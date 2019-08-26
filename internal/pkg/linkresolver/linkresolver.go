package linkresolver

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"regexp"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type (
	Link struct {
		Title       string
		Description string
		Image       struct {
			URL string
		}
	}
	Resolver struct {
		client *http.Client
	}

	attrFilterFunc func(name string) bool
)

// New return a new resolver instance
func New() *Resolver {
	c := http.DefaultClient
	c.Timeout = 5 * time.Second
	c.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	return &Resolver{
		client: c,
	}
}

// Resolve resolve the url information
func (r *Resolver) Resolve(url string) (*Link, error) {
	res, err := r.client.Get(url)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("got status=%d, want status=%d", res.StatusCode, http.StatusOK)
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}
	link := &Link{}
	doc.Find("head").Each(func(i int, e *goquery.Selection) {
		e.Find("title").Each(func(i int, e *goquery.Selection) {
			link.Title = e.Text()
		})
		e.Find("meta").Each(func(i int, e *goquery.Selection) {
			if description, ok := attrVal(e, "name", matchNames("description"), "content"); ok {
				link.Description = description
			}
			if image, ok := attrVal(e, "name", matchNames("image.thumb", "image", "icon"), "content"); ok && isImageLink(image) {
				link.Image.URL = image
			}
		})
		// try to get image if not found in meta
		if link.Image.URL != "" {
			return
		}
		e.Find("meta").Each(func(i int, e *goquery.Selection) {
			if image, ok := attrVal(e, "property", matchNames("image.thump", "image", "icon"), "content"); ok && isImageLink(image) {
				link.Image.URL = image
			}
		})
		if link.Image.URL != "" {
			return
		}
		e.Find("link").Each(func(i int, e *goquery.Selection) {
			if image, ok := attrVal(e, "rel", matchNames("image", "icon"), "href"); ok && isImageLink(image) {
				link.Image.URL = image
			}
		})
	})
	if link.Description != "" {
		return link, nil
	}
	// description is still empty, try to read the body
	// TODO implement me
	return link, nil
}

// attrVal try to lookup the attribute by lookupName and value of the valuteAttrName
func attrVal(e *goquery.Selection, lookupName string, filter attrFilterFunc, valueAttrName string) (string, bool) {
	if name, ok := e.Attr(lookupName); ok && filter(name) {
		if v, ok := e.Attr(valueAttrName); ok && v != "" {
			return v, true
		}
	}
	return "", false
}

// matchNames return AttrFilterFunc by matching the given names by order
func matchNames(patterns ...string) attrFilterFunc {
	return func(name string) bool {
		for _, p := range patterns {
			if matched, _ := regexp.MatchString(p, name); matched {
				return true
			}
		}
		return false
	}
}
