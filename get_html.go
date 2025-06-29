package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getHTML(rawURL string) (string, error) {
	res, err := http.Get(rawURL)
	if err != nil {
		return "", fmt.Errorf("error creating request: %w", err)
	}
	defer res.Body.Close()
	if res.StatusCode > 399 {
		return "", fmt.Errorf("request faied with status code: %d", res.StatusCode)
	}
	header := res.Header.Get("content-type")
	if !strings.HasPrefix(header, "text/html") {
		return "", fmt.Errorf("incompatible content type: %s", header)
	}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
