package mt

import (
	"net"
	"time"
)

// Comment represents one comment on the entry.
type Comment struct {
	// Author is the name of the author of the comment.
	Author string
	// Date is the date on which the comment was posted.
	Date time.Time
	// URL is the URL of the author of the content.
	URL string
	// IP is the IP address of the author of the comment.
	IP net.IP
	// Email is the email address of the author of the comment.
	Email string
	Body  string
}
