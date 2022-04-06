package main

import (
	"flag"
	"fmt"
	"github/hawk911/pkg/crawler"
	"github/hawk911/pkg/crawler/spider"
	"strings"
)

func main() {
	var total []crawler.Document

	serv := spider.New()
	urls := [2]string{"https://golang.com", "https://go.dev"}

	for _, url := range urls {
		data, err := serv.Scan(url, 2)
		if err != nil {
			fmt.Println("Scan error. Please, try again.")
			continue
		}
		total = append(total, data...)
	}

	sWord := flag.String("s", "", "Please, input word(s)")
	flag.Parse()

	if *sWord == "" {
		fmt.Println("Please, read help: -help \n Thanks.")
		return
	}

	var fData []crawler.Document

	for _, v := range total {
		if strings.Contains(strings.ToLower(v.Title), strings.ToLower(*sWord)) {
			fData = append(fData, v)
		}
	}

	if len(fData) == 0 {
		fmt.Println("Nothing was found.")
		return
	}

	for _, v := range fData {
		fmt.Printf("Title: %s\nURL: %s\n\n", v.Title, v.URL)
	}
	fmt.Println("Thanks for using our solution!")
}
