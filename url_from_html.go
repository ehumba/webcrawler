package main

import (
	"errors"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func getURLsFromHTML(htmlBody string, BaseURL *url.URL) ([]string, error) {

	reader := strings.NewReader(htmlBody)
	tree, err := html.Parse(reader)
	if err != nil {
		return nil, errors.New("could not parse html body")
	}
	var newURLs []string
	var crawler func(*html.Node)
	crawler = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "a" {
			for _, att := range node.Attr {
				if att.Key == "href" {
					parsedURLRef, err := url.Parse(att.Val)
					if err != nil {
						continue
					}
					finalURL := BaseURL.ResolveReference(parsedURLRef)
					newURLs = append(newURLs, finalURL.String())
				}
			}
		}
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			crawler(child)
		}

	}
	crawler(tree)
	return newURLs, nil
}
