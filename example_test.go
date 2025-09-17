package mt_test

import (
	"fmt"
	"strings"

	"github.com/aereal/mt"
)

func ExampleParse() {
	r := strings.NewReader(exampleImportFile)
	fmt.Println("entries:")
	for entry, err := range mt.Parse(r) {
		if err != nil {
			panic(err)
		}
		fmt.Printf("title: %s", entry.Title)
	}
	// Output:
	// entries:
	// title: A dummy title
}

var exampleImportFile = `
AUTHOR: Foo Bar
TITLE: A dummy title
BASENAME: filename
STATUS: Publish
ALLOW COMMENTS: 1
ALLOW PINGS: 1
CONVERT BREAKS: richtext
PRIMARY CATEGORY: News
CATEGORY: News
CATEGORY: Product
DATE: 08/08/2007 03:00:00 PM
TAGS: "Movable Type",foo,bar
-----
BODY:
これは本文です。
-----
EXTENDED BODY:
ここに追記の本文が表示されます。
-----
COMMENT:
AUTHOR: Foo
DATE: 01/31/2002 15:47:06
ここに
このコメントの本文が来ます。
-----
COMMENT:
AUTHOR: Bar
DATE: 02/01/2002 04:02:07 AM
IP: 205.66.1.32
EMAIL: me@bar.com
これは2番目の
コメントです。 これは
ここまで来ます。
-----
--------
`
