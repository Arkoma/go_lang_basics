// Get content-type from sites
package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func returnType(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("error: %s/n", err)
		return
	}
	defer resp.Body.Close()
	ctype := resp.Header.Get("content-type")
	fmt.Printf("%s -> %s\n", url, ctype)
}

func siteSerial(urls []string) {
	for _, url := range urls {
		returnType(url)
	}
}

func siteConcurrent(urls []string) {
	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1) // tells how many go routines the wait group needs to wait for
		go func(url string) {
			returnType(url)
			wg.Done()
		}(url) // this go func {} is defining a function the () after the closing '}' is for calling said function immediately with parameters if they exist
		wg.Wait()
	}
}

func main() {
	urls := []string{
		"https://golang.org",
		"https://api.gitub.com",
		"https://httpbin.org/ip",
	}
	start := time.Now()
	siteSerial(urls)
	fmt.Println(time.Since(start))

	start2 := time.Now()
	siteConcurrent(urls)
	fmt.Println(time.Since(start2))
}
