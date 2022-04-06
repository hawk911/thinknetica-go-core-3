package main

import (
	"flag"
	"fmt"
	"github/hawk911/pkg/crawler"
	"github/hawk911/pkg/crawler/spider"
	"github/hawk911/pkg/index"
	"strings"
)

var urls = []string{"https://golang.com", "https://go.dev"}

func scan(urls []string, idx *index.Service, serv *spider.Service) ([]crawler.Document, error) {
	var total []crawler.Document
	var count = 0

	for _, url := range urls {
		data, err := serv.Scan(url, 2)
		if err != nil {
			fmt.Println("Scan error. Please, try again.")
			return total, err
		}

		for _, i := range data {
			i.ID = count
			idx.Add(i.Title, i.ID)
			total = append(total, i)
			count++
		}
	}
	return total, nil
}
func main() {

	serv := spider.New()
	idx := index.New()
	t, err := scan(urls, idx, serv)
	if err != nil {
		fmt.Printf("Error is %s\n", err)
	}

	sWord := flag.String("s", "", "Please, input word(s)")
	flag.Parse()

	if *sWord == "" {
		fmt.Println("Please, read help: -help \n Thanks.")
		return
	}

	ids := idx.Search(strings.ToLower(*sWord))

	if len(ids) == 0 {
		fmt.Println("Nothing was found.")
		return
	}

	for _, id := range ids {
		fmt.Printf("Title: %s\nURL: %s\n\n", t[id].Title, t[id].URL)
	}
	fmt.Println("Thanks for using our solution!")
}
