package main

import (
	"fmt"
	"sort"
)

func printReport(pages map[string]int, baseURL string) {
	fmt.Println("=============================")
	fmt.Printf("REPORT for %s\n", baseURL)
	fmt.Println("=============================")

	type kv struct {
		Key   string
		Value int
	}

	var pagesSlice []kv
	for page, count := range pages {
		pagesSlice = append(pagesSlice, kv{page, count})
	}

	sort.Slice(pagesSlice, func(i, j int) bool {
		return pagesSlice[i].Value > pagesSlice[j].Value
	})

	for _, pageCount := range pagesSlice {
		if pageCount.Value > 1 {
			fmt.Printf("Found %v internal links to %s\n", pageCount.Value, pageCount.Key)
		} else {
			fmt.Printf("Found %v internal link to %s\n", pageCount.Value, pageCount.Key)
		}
	}
}
