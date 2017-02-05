package formathipchat

import (
	"errors"
	"fmt"
	"golang.org/x/net/html"
	"strings"
)

type Message struct {
	Color          string `json:"color"`
	Message        string `json:"message"`
	Message_format string `json:"message_format"`
	Notify         bool   `json:"notify"`
}

func checkTagAttr(t html.Token, tag string, compare string) bool {
	for _, attr := range t.Attr {
		if attr.Key == tag {
			return attr.Val == compare
		}
	}
	return false
}

func getTagAttr(t html.Token, tag string) (val string, err error) {
	for _, attr := range t.Attr {
		if attr.Key == tag {
			val = attr.Val
			return
		}
	}
	err = errors.New(fmt.Sprintf("failed to parse %s", tag))
	return
}

func parseHtml(data []byte) (title string, src string, err error) {
	z := html.NewTokenizer(strings.NewReader(string(data)))

	for {
		tt := z.Next()

		switch {
		case tt == html.ErrorToken:
			// End of the document
			return
		case tt == html.SelfClosingTagToken:
			t := z.Token()

			isComicImg := t.Data == "img" && checkTagAttr(t, "class", "img-responsive img-comic")

			// Hack: first img tag is the latest comic
			if isComicImg {
				title, err = getTagAttr(t, "alt")
				if err != nil {
					return
				}
				src, err = getTagAttr(t, "src")
				if err != nil {
					return
				}
				return
			}
		}
	}
	err = errors.New(fmt.Sprintf("failed to find img tag"))
	return
}

func Format(data []byte) (Message, error) {
	title, src, err := parseHtml(data)
	if err != nil {
		return Message{}, err
	}

	msg := Message{
		Color:          "gray",
		Message:        "<span>" + title + "</span><br><img src='" + src + "'/>",
		Message_format: "html",
		Notify:         true,
	}

	return msg, nil
}
