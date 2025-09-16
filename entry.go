package mt

import "time"

type Entry struct {
	// Author is the author of the entry.
	//
	// refs. https://movabletype.org/documentation/appendices/import-export-format.html#author
	Author string

	// Title is the title of the entry.
	//
	// refs. https://movabletype.org/documentation/appendices/import-export-format.html#title
	Title string

	// Basename is the basename of the entry.
	//
	// refs. https://movabletype.org/documentation/appendices/import-export-format.html#basename
	Basename string

	// Date is the authored-on or published date of the entry.
	//
	// refs. https://movabletype.org/documentation/appendices/import-export-format.html#date
	Date time.Time

	// PrimaryCategory is the primary category to which the entry is assigned.
	//
	// refs. https://movabletype.org/documentation/appendices/import-export-format.html#primary-category
	PrimaryCategory string

	// Category is a secondary category to which the entry is assigned.
	//
	// refs. https://movabletype.org/documentation/appendices/import-export-format.html#category
	Category []string

	// Tags is the tags associated with an entry.
	//
	// refs. https://movabletype.org/documentation/appendices/import-export-format.html#tags
	Tags []string

	// Status is the post status of the entry.
	//
	// refs. https://movabletype.org/documentation/appendices/import-export-format.html#status
	Status string

	// AllowComments is the value for the "allow comments" flag for the entry.
	//
	// refs. https://movabletype.org/documentation/appendices/import-export-format.html#allow-comments
	AllowComments bool

	// AllowPings is the value for the "allow pings" flag for the entry.
	//
	// refs. https://movabletype.org/documentation/appendices/import-export-format.html#allow-pings
	AllowPings bool

	// ConvertBreaks is the value for the "convert breaks" flag for the entry.
	//
	// refs. https://movabletype.org/documentation/appendices/import-export-format.html#convert-breaks
	ConvertBreaks ConvertBreaks

	// Body is the body of the entry.
	Body string

	// ExtendedBody is the extended body of the entry.
	ExtendedBody string

	// Excerpt is the excerpt of the entry.
	Excerpt string

	Comments []*Comment
}
