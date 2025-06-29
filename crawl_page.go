package main

import (
	"fmt"
	"net/url"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
	parsedBaseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		fmt.Println("could not parse base url")
		return
	}
	parsedCurrentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Println("could not parse current url")
		return
	}
	if parsedBaseURL.Hostname() != parsedCurrentURL.Hostname() {
		return
	}

	currentNormalized, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("could not normalize url: %v", err)
		return
	}

	pages[currentNormalized] += 1
	if pages[currentNormalized] > 1 {
		return
	}

	html, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("failed to fetch html: %v", err)
		return
	}
	fmt.Print(html)

	urls, err := getURLsFromHTML(html, rawBaseURL)
	if err != nil {
		fmt.Printf("failed to fetch urls from html body: %v", err)
		return
	}

	for _, url := range urls {
		crawlPage(rawBaseURL, url, pages)
	}

}
