package mt

import (
	"net"
	"time"
)

type Comment struct {
	Author string
	Date   time.Time
	URL    string
	IP     net.IP
	Email  string
	Body   string
}
