package main

import (
	"fmt"
	"net/url"
	"os"
	"sync"
)

func main() {
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) < 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	}

	if len(argsWithoutProg) > 1 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	URLInput := argsWithProg[1]
	parsedInput, err := url.Parse(URLInput)
	if err != nil {
		os.Exit(1)
	}

	fmt.Printf("starting crawl of: %v\n", URLInput)

	cfg := &config{
		pages:              make(map[string]int),
		baseURL:            parsedInput,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, 5),
		wg:                 &sync.WaitGroup{},
	}

	cfg.crawlPage(URLInput)
	cfg.wg.Wait()
	fmt.Println(cfg.pages)
}
