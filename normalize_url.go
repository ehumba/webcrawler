package main

import (
	"net/url"
	"strings"
)

func normalizeURL(rawURL string) (string, error) {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}
	parsedURL.Fragment = ""
	parsedURL.RawQuery = ""
	hostAndPath := parsedURL.Host + parsedURL.Path
	normalized := strings.Trim(hostAndPath, "/")

	return normalized, nil
}
