// Copyright (c) 2023, The GoKi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gidom

import (
	"slices"

	"goki.dev/gi/v2/gi"
	"golang.org/x/net/html"
)

// ExtractText recursively extracts all of the text from the children
// of the given [*html.Node], adding any appropriate inline markup for
// formatted text. It adds any non-text elements to the given [gi.Widget]
// using [ReadHTMLNode]. It should not be called on text nodes themselves;
// for that, you can directly access the [html.Node.Data] field.
func ExtractText(par gi.Widget, n *html.Node) string {
	if n.FirstChild == nil {
		return ""
	}
	return extractTextImpl(par, n.FirstChild)
}

func extractTextImpl(par gi.Widget, n *html.Node) string {
	str := ""
	if n.Type == html.TextNode {
		str += n.Data
	}
	it := IsText(n)
	if !it {
		ReadHTMLNode(par, n)
	}
	if it && n.FirstChild != nil {
		if n.Type == html.ElementNode {
			tag := n.DataAtom.String()
			str += "<" + tag + ">"
		}
		str += extractTextImpl(par, n.FirstChild)
		if n.Type == html.ElementNode {
			tag := n.DataAtom.String()
			str += "</" + tag + ">"
		}
	}
	if n.NextSibling != nil {
		str += extractTextImpl(par, n.NextSibling)
	}
	return str
}

// TextTags are all of the node tags that result in a true return value for [IsText].
var TextTags = []string{
	"a", "abbr", "b", "bdi", "bdo", "br", "cite", "code", "data", "dfn",
	"em", "i", "kbd", "mark", "q", "rp", "rt", "ruby", "s", "samp", "small",
	"span", "strong", "sub", "sup", "time", "u", "var", "wbr",
}

// IsText returns true if the given node is a [html.TextNode] or
// an [html.ElementNode] designed for holding text (p, span, b, code, etc),
// and false otherwise.
func IsText(n *html.Node) bool {
	if n.Type == html.TextNode {
		return true
	}
	if n.Type == html.ElementNode {
		tag := n.DataAtom.String()
		return slices.Contains(TextTags, tag)
	}
	return false
}
