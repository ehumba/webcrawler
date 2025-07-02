package main

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"sync"
)

func main() {
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) < 3 {
		fmt.Println("not enough arguments provided")
		os.Exit(1)
	}

	if len(argsWithoutProg) > 3 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	URLInput := argsWithProg[1]
	parsedInput, err := url.Parse(URLInput)
	if err != nil {
		os.Exit(1)
	}

	maxConcurrency, err := strconv.Atoi(argsWithProg[2])
	if err != nil {
		os.Exit(1)
	}
	maxPagesInput, err := strconv.Atoi(argsWithProg[3])
	if err != nil {
		os.Exit(1)
	}

	cfg := &config{
		pages:              make(map[string]int),
		baseURL:            parsedInput,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, maxConcurrency),
		wg:                 &sync.WaitGroup{},
		maxPages:           maxPagesInput,
	}

	cfg.crawlPage(URLInput)
	cfg.wg.Wait()
	printReport(cfg.pages, URLInput)
}
