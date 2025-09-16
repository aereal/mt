package mt_test

import (
	"embed"
	"net"
	"testing"
	"time"

	"github.com/aereal/mt"
	"github.com/google/go-cmp/cmp"
)

//go:embed testdata/*
var testdata embed.FS

func TestParse(t *testing.T) {
	t.Parallel()

	f, err := testdata.Open("testdata/ok.txt")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = f.Close() })
	gotEntries := make([]*mt.Entry, 0)
	for entry, err := range mt.Parse(f) {
		if err != nil {
			t.Fatal(err)
		}
		gotEntries = append(gotEntries, entry)
	}
	want := []*mt.Entry{
		{
			Author:          "Foo Bar",
			Title:           "A dummy title",
			Basename:        "filename",
			Status:          "Publish",
			Date:            time.Date(2007, time.August, 8, 15, 0, 0, 0, time.Local),
			PrimaryCategory: "News",
			Category:        []string{"News", "Product"},
			Tags:            []string{"Movable Type", "foo", "bar"},
			AllowComments:   true,
			AllowPings:      true,
			ConvertBreaks:   mt.ConvertBreaksRichtext,
			Body:            "これは本文です。\n",
			ExtendedBody:    "ここに追記の本文が表示されます。\n",
			Comments: []*mt.Comment{
				{
					Author: "Foo",
					Date:   time.Date(2002, time.January, 31, 15, 47, 6, 0, time.Local),
					Body:   "ここに\nこのコメントの本文が来ます。\n",
				},
				{
					Author: "Bar",
					Date:   time.Date(2002, time.February, 1, 4, 2, 7, 0, time.Local),
					Body:   "これは2番目の\nコメントです。 これは\nここまで来ます。\n",
					IP:     net.IPv4(205, 66, 1, 32),
					Email:  "me@bar.com",
				},
			},
		},
		{
			Author:          "Foo Bar",
			Title:           "2件目の記事 ",
			Basename:        "filename",
			Status:          "Publish",
			Date:            time.Date(2007, time.August, 8, 15, 0, 0, 0, time.Local),
			PrimaryCategory: "News",
			Category:        []string{"News", "Product"},
			Tags:            []string{"Movable Type", "foo", "bar"},
			AllowComments:   true,
			AllowPings:      true,
			ConvertBreaks:   mt.ConvertBreaksRichtext,
			Body:            "これは2番目の記事の本文です。 これは\n複数行から成ります。\n",
			Excerpt:         "この記事は追記がありませんが、\n概要はあります。 特殊な例です。\n",
			Comments: []*mt.Comment{
				{
					Author: "Quux",
					Date:   time.Date(2002, time.January, 31, 16, 23, 1, 0, time.Local),
					URL:    `<a href="http://www.quux.com/">http://www.quux.com/</a>`,
					Body:   "この記事に対する最初のコメントを示します。\n",
				},
			},
		},
	}
	if diff := cmp.Diff(want, gotEntries); diff != "" {
		t.Errorf("(-want, +got):\n%s", diff)
	}
}
