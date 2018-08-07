package main

import (
	"fmt"
	"net/http"
)

type sites []string

func main() {
	siteList := sites{
		"https://www.google.com",
		"https://www.amazon.com",
		"https://www.youtube.com",
		"https://www.hahahlaksdf.com",
		"https://www.facebook.com",
	}

	c := make(chan string)

	for _, link := range siteList {
		go checker(link, c)
	}

	for {
		go checker(<-c, c)
	}

}

func checker(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "Status: [DOWN]")
		c <- link
		return
	}

	fmt.Println(link, "Status: [UP]")
	c <- link
}
