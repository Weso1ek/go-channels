package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://go.dev",
		"http://amazon.com",
	}

	// channel
	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	//fmt.Println(<-c)
	//fmt.Println(<-c)
	//fmt.Println(<-c)
	//fmt.Println(<-c)
	//fmt.Println(<-c)

	//for i := 0; i < len(links); i++ {
	for l := range c {
		//time.Sleep(5 * time.Second)
		//go checkLink(l, c)
		//fmt.Println(<-c)
		go func() { // anonymous function
			time.Sleep(5 * time.Second)
			checkLink(l, c)
		}()
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "is down")
		c <- link
		return
	}

	fmt.Println(link, "is up")
	c <- link
}
