package main

import (
	"net/url"
	"sync"
)

type config struct {
	pages              map[string]int
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
	maxPages           int
}

func (cfg *config) addPageVisit(normalizedURL string) (isFirst bool) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	_, ok := cfg.pages[normalizedURL]
	if ok {
		return false
	} else {
		cfg.pages[normalizedURL] = 1
		return true
	}
}
