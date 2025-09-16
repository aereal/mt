package mt

import (
	"bytes"
	"errors"
	"fmt"
)

type ParseError struct {
	Text    string
	LineNum int
	Field   Field
	Err     error
}

var _ error = (*ParseError)(nil)

func (e *ParseError) Error() string {
	buf := new(bytes.Buffer)
	fmt.Fprintf(buf, "parse failed: line=%d", e.LineNum)
	if e.Field != "" {
		fmt.Fprintf(buf, " field=%s", e.Field)
	}
	fmt.Fprintf(buf, " error=%s text=%q", e.Err, e.Text)
	return buf.String()
}

func (e *ParseError) Unwrap() error { return e.Err }

func (e *ParseError) Is(target error) bool {
	pErr := new(ParseError)
	if !errors.As(target, &pErr) {
		return false
	}
	return e.LineNum == pErr.LineNum && e.Field == pErr.Field && errors.Is(e.Err, pErr.Err)
}

var ErrNoKeyValueDelimiterFound NoKeyValueDelimiterFoundError

type NoKeyValueDelimiterFoundError struct{}

var _ error = NoKeyValueDelimiterFoundError{}

func (NoKeyValueDelimiterFoundError) Error() string { return "no key-value delimiter found" }

type UnexpectedBooleanNumberError struct {
	Actual int
}

var _ error = (*UnexpectedBooleanNumberError)(nil)

func (e *UnexpectedBooleanNumberError) Error() string {
	return fmt.Sprintf("expected 0 or 1 but got %d", e.Actual)
}

type InvalidConvertBreaksError struct {
	Value string
}

var _ error = (*InvalidConvertBreaksError)(nil)

func (e *InvalidConvertBreaksError) Error() string {
	return fmt.Sprintf("invalid convert breaks: %q", e.Value)
}
