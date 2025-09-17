package mt

import (
	"bufio"
	"fmt"
	"io"
	"iter"
	"net"
	"strconv"
	"strings"
	"time"
	"unicode"
)

const (
	keyValueDelimiter       = ":"
	multilineFieldDelimiter = "-----"
	entryDelimiter          = "--------"
	tagDelimiter            = ","
	tagQuotation            = `"`
)

func Parse(r io.Reader, opts ...ParseOption) iter.Seq2[*Entry, error] {
	return func(yield func(*Entry, error) bool) {
		cfg := new(parseConfig)
		for _, o := range opts {
			o.applyParseOption(cfg)
		}
		if cfg.loc == nil {
			cfg.loc = time.Local
		}

		scanner := bufio.NewScanner(r)
		scanner.Split(bufio.ScanLines)

		var text string
		var lineNum int
		reportParseError := func(field Field, err error) {
			_ = yield(nil, &ParseError{LineNum: lineNum, Field: field, Text: text, Err: err})
		}

		current := &Entry{
			ConvertBreaks: cfg.defaultConvertBreaks,
			AllowComments: cfg.defaultAllowComments,
		}
		for scanner.Scan() {
			lineNum++
			text = scanner.Text()
			if text == multilineFieldDelimiter {
				continue
			}
			if text == entryDelimiter {
				if !yield(current, nil) {
					return
				}
				current = &Entry{}
				continue
			}
			key, value, ok := strings.Cut(text, keyValueDelimiter)
			if !ok {
				reportParseError(FieldNone, ErrNoKeyValueDelimiterFound)
				return
			}
			value = strings.TrimLeftFunc(value, unicode.IsSpace)

			switch Field(key) {
			case FieldBody:
				for scanner.Scan() {
					lineNum++
					l := scanner.Text()
					if l == multilineFieldDelimiter {
						break
					}
					current.Body += l + "\n"
				}
			case FieldExtendedBody:
				for scanner.Scan() {
					lineNum++
					l := scanner.Text()
					if l == multilineFieldDelimiter {
						break
					}
					current.ExtendedBody += l + "\n"
				}
			case FieldExcerpt:
				for scanner.Scan() {
					lineNum++
					l := scanner.Text()
					if l == multilineFieldDelimiter {
						break
					}
					current.Excerpt += l + "\n"
				}
			case FieldComment:
				for comment, err := range parseComments(&lineNum, reportParseError, cfg.loc, scanner) {
					if err != nil {
						reportParseError(FieldComment, fmt.Errorf("parseComments: %w", err))
						return
					}
					current.Comments = append(current.Comments, comment)
				}
			case FieldAuthor:
				current.Author = value
			case FieldTitle:
				current.Title = value
			case FieldBasename:
				current.Basename = value
			case FieldStatus:
				st, err := ParseStatus(value)
				if err != nil {
					reportParseError(FieldStatus, err)
					return
				}
				current.Status = st
			case FieldAllowComments:
				v, err := parseIntBool(value)
				if err != nil {
					reportParseError(FieldAllowComments, err)
					return
				}
				current.AllowComments = v
			case FieldAllowPings:
				v, err := parseIntBool(value)
				if err != nil {
					reportParseError(FieldAllowPings, err)
					return
				}
				current.AllowPings = v
			case FieldConvertBreaks:
				cb, err := ParseConvertBreaks(value)
				if err != nil {
					reportParseError(FieldConvertBreaks, err)
					return
				}
				current.ConvertBreaks = cb
			case FieldPrimaryCategory:
				current.PrimaryCategory = value
			case FieldCategory:
				current.Category = append(current.Category, value)
			case FieldTags:
				tags, err := parseTags(value)
				if err != nil {
					reportParseError(FieldTags, err)
					return
				}
				current.Tags = append(current.Tags, tags...)
			case FieldDate:
				var err error
				current.Date, err = parseDate(value, cfg.loc)
				if err != nil {
					reportParseError(FieldDate, err)
					return
				}
			default:
				// noop
			}
		}
	}
}

func parseComments(lineNum *int, reportParseError func(Field, error), loc *time.Location, scanner *bufio.Scanner) iter.Seq2[*Comment, error] {
	return func(yield func(*Comment, error) bool) {
		var text string
		current := &Comment{}
		for scanner.Scan() {
			*lineNum++
			text = scanner.Text()
			if text == multilineFieldDelimiter {
				continue
			}
			key, value, ok := strings.Cut(text, keyValueDelimiter)
			if !ok { // the body may be appeared if no metadata found
				current.Body += text + "\n"

				for scanner.Scan() {
					*lineNum++
					l := scanner.Text()
					if l == multilineFieldDelimiter {
						break
					}
					current.Body += l + "\n"
				}
				if !yield(current, nil) {
					return
				}
				current = &Comment{}
				break
			}
			value = strings.TrimLeftFunc(value, unicode.IsSpace)

			switch Field(key) {
			case FieldAuthor:
				current.Author = value
			case FieldEmail:
				current.Email = value
			case FieldURL:
				current.URL = value
			case FieldDate:
				var err error
				current.Date, err = parseDate(value, loc)
				if err != nil {
					reportParseError(FieldDate, err)
					return
				}
			case FieldIP:
				current.IP = net.ParseIP(value)
			default:
				// noop
			}
		}
	}
}

var (
	dateLayout   = "01/02/2006 15:04:05"
	dateWithNoon = "01/02/2006 03:04:05 PM"
)

func parseDate(s string, loc *time.Location) (time.Time, error) {
	layout := dateLayout
	if strings.HasSuffix(s, " AM") || strings.HasSuffix(s, " PM") {
		layout = dateWithNoon
	}
	parsed, err := time.ParseInLocation(layout, s, loc)
	if err != nil {
		return time.Time{}, err
	}
	return parsed, nil
}

func parseTags(s string) ([]string, error) {
	ret := make([]string, 0)
	for v := range strings.SplitSeq(s, tagDelimiter) {
		if strings.HasPrefix(v, tagQuotation) {
			var err error
			v, err = strconv.Unquote(v)
			if err != nil {
				return nil, err
			}
		}
		ret = append(ret, v)
	}
	return ret, nil
}

func parseIntBool(s string) (bool, error) {
	v, err := strconv.Atoi(s)
	if err != nil {
		return false, err
	}
	switch v {
	case 0:
		return false, nil
	case 1:
		return true, nil
	default:
		return false, &UnexpectedBooleanNumberError{Actual: v}
	}
}
