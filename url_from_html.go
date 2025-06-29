package main

import (
	"errors"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	// Need to read through the HTML body, add the path to the raw base url, and add it to the
	// array, or add the whole url if it is an absolute url.
	// Create an io.Reader with strings.NewReader(htmlBody). Then use this reader with html.Parse
	// to create a tree of nodes. Navigate it and find the <a anchors (the Node.Data?)
	// If it is an absolut URL (starts with https://), just add it to the slice as is, if not, append it to the rawBaseURL
	// and add to the slice.
	parsedBaseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return nil, errors.New("could not parse base url")
	}
	reader := strings.NewReader(htmlBody)
	tree, err := html.Parse(reader)
	if err != nil {
		return nil, errors.New("could not parse html body")
	}
	var newURLs []string
	for node := range tree.Descendants() {
		if node.Type == html.ElementNode && node.Data == "a" {
			for _, att := range node.Attr {
				if att.Key == "href" {
					parsedURLRef, err := url.Parse(att.Val)
					if err != nil {
						return nil, errors.New("could not parse url reference")
					}
					finalURL := parsedBaseURL.ResolveReference(parsedURLRef)
					newURLs = append(newURLs, finalURL.String())
				}
			}
		}
	}
	return newURLs, nil
}
