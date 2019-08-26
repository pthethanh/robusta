package htmlutil

import (
	"io"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

type (
	Link struct {
		Href string
	}
)

// ReadHeaderLinks return list of links in the header
func ReadHeaderLinks(reader io.Reader) ([]Link, error) {
	links := make([]Link, 0)
	doc, err := html.Parse(reader)
	if err != nil {
		return nil, err
	}
	head := findNode(doc, atom.Head)
	if head == nil {
		return links, nil
	}
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.DataAtom == atom.Link {
			for _, attr := range n.Attr {
				if attr.Key == atom.Href.String() {
					links = append(links, Link{
						Href: attr.Val,
					})
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(head)
	return links, nil
}

func findNode(n *html.Node, t atom.Atom) *html.Node {
	var found *html.Node
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.DataAtom == t {
			found = n
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(n)
	return found
}
