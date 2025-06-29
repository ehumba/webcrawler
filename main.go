package main

import (
	"fmt"
	"os"
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

	url := argsWithProg[1]

	fmt.Printf("starting crawl of: %v\n", url)
	pages := make(map[string]int)
	crawlPage(url, url, pages)
	fmt.Print(pages)
}
