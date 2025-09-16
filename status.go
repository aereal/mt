package mt

import "strings"

const (
	StatusDraft   Status = "draft"
	StatusPublish Status = "publish"
	StatusFuture  Status = "future"
)

type Status string

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
