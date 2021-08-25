package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{"http://google.com", "http://facebook.com", "http://stackoverflow.com", "http://golang.org", "http://amazon.com"}
	c := make(chan string)
	for _, l := range links {
		go checkUrl(l, c)
	}

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkUrl(link, c)
		}(l)
	}
}
func checkUrl(url string, c chan string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Println(url, "might be down")
		c <- url
		return
	}
	fmt.Println(url, "is up")
	c <- url
}
