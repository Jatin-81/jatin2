package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

type lang struct {
	name  string
	url   string
	bytes int64
	time  float64
}

func (l lang) String() string {
	return fmt.Sprintf("Name: %v\n Url: %v\n Bytes : %v\n Time Taken: %v", l.name, l.url, l.bytes, l.time)
}

func printDetails(lang *lang) {
	fmt.Printf("\n %v\n", lang)
}

func crawl(pfunc func(lang *lang), lang *lang) {
	start := time.Now()
	resp, err := http.Get(lang.url)

	if err != nil {
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)

	if err != nil {
		fmt.Printf("while reading %s: %v", lang.url, err)
		return
	}

	defer resp.Body.Close()

	lang.bytes = nbytes
	lang.time = time.Since(start).Seconds()

	pfunc(lang)
}

func main() {

	urlMap := map[string]string{
		"Python": "https://www.python.org/",
		"Ruby":   "https://www.ruby-lang.org/en/",
		"Golang": "https://golang.org/",
	}

	startTime := time.Now()

	for key, value := range urlMap {
		input := lang{
			name: key,
			url:  value,
		}
		crawl(printDetails, &input)
	}

	fmt.Println("\n Total time to crawl data from sites is : ", time.Since(startTime).Seconds())
}
