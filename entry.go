package mt

import "time"

type Entry struct {
	// Author is the author of the entry.
	Author string

	// Title is the title of the entry.
	Title string

	// Basename is the basename of the entry.
	Basename string

	// Date is the authored-on or published date of the entry.
	Date time.Time

	// PrimaryCategory is the primary category to which the entry is assigned.
	PrimaryCategory string

	// Category is a secondary category to which the entry is assigned.
	Category []string

	// Tags is the tags associated with an entry.
	Tags []string

	// Status is the post status of the entry.
	Status Status

	// AllowComments is the value for the "allow comments" flag for the entry.
	AllowComments bool

	// AllowPings is the value for the "allow pings" flag for the entry.
	AllowPings bool

	// ConvertBreaks is the value for the "convert breaks" flag for the entry.
	ConvertBreaks ConvertBreaks

	// Body is the body of the entry.
	Body string

	// ExtendedBody is the extended body of the entry.
	ExtendedBody string

	// Excerpt is the excerpt of the entry.
	Excerpt string

	Comments []*Comment
}
