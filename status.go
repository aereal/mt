package mt

import "strings"

const (
	StatusDraft   Status = "draft"
	StatusPublish Status = "publish"
	StatusFuture  Status = "future"
)

// Status is the post status of the entry.
type Status string

// ParseStatus returns a [Status] corresponding to the argument.
func ParseStatus(v string) (Status, error) {
	switch s := Status(strings.ToLower(v)); s {
	case StatusDraft,
		StatusPublish,
		StatusFuture:
		return s, nil
	default:
		return "", &InvalidStatusError{Value: v}
	}
}
