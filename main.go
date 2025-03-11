package main

import "fmt"

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://go.dev",
		"http://amazon.com",
	}

	for _, link := range links {
		fmt.Println(link)
	}

	fmt.Println(links)
}
