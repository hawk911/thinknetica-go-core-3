package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github/hawk911/pkg/crawler"
	"github/hawk911/pkg/crawler/spider"
	"github/hawk911/pkg/index"
	"io"
	"os"
	"strings"
)

var urls = []string{"https://golang.com", "https://go.dev"}

func get(r io.Reader) ([]byte, error) {
	return io.ReadAll(r)
}

func store(w io.Writer, b []byte) error {
	_, err := w.Write(b)
	return err
}

func scan(urls []string, idx *index.Service, serv *spider.Service) ([]crawler.Document, error) {
	var total []crawler.Document
	var data []crawler.Document
	if _, err := os.Stat("./archive.json"); err == nil {
		fmt.Println("File exists.")

		f, err := os.Open("./archive.json")
		defer f.Close()
		if err != nil {
			fmt.Println("Opening is error. Please, try again.")
			return nil, err
		}
		b, _ := get(f)

		err = json.Unmarshal(b, &data)
		if err != nil {
			fmt.Println("JSON Reading is error. Please, try again.")
			return nil, err
		}

	} else {
		fmt.Println("File does not exist. Web spider is starting...")

		for _, url := range urls {
			data, err = serv.Scan(url, 2)
			if err != nil {
				fmt.Println("Scan error. Please, try again.")
				return nil, err
			}
		}

		content, err := json.Marshal(data)
		if err != nil {
			fmt.Println("error:", err)
		}

		f, err := os.Create("./archive.json")
		if err != nil {
			fmt.Println("Creating is error. Please, try again.")
			return nil, err
		}
		defer f.Close()
		err = store(f, content)

	}

	var count = 0
	for _, i := range data {
		i.ID = count
		idx.Add(i.Title, i.ID)
		total = append(total, i)
		count++
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
