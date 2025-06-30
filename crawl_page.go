package main

import (
	"fmt"
	"net/url"
)

func (cfg *config) crawlPage(rawCurrentURL string) {
	parsedCurrentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Println("could not parse current url")
		return
	}
	if cfg.baseURL.Hostname() != parsedCurrentURL.Hostname() {
		return
	}

	currentNormalized, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("could not normalize url: %v", err)
		return
	}

	// Check if a page was already visited and return if yes.
	if cfg.addPageVisit(currentNormalized) == false {
		return
	}

	html, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("failed to fetch html: %v", err)
		return
	}
	fmt.Print(html)

	urls, err := getURLsFromHTML(html, cfg.baseURL)
	if err != nil {
		fmt.Printf("failed to fetch urls from html body: %v", err)
		return
	}

	for _, url := range urls {
		cfg.wg.Add(1)
		go func(link string) {
			cfg.concurrencyControl <- struct{}{}
			defer cfg.wg.Done()
			defer func() { <-cfg.concurrencyControl }() // release token
			cfg.crawlPage(url)
		}()

	}

}
