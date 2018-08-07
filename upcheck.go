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

	for i := 0; i < len(siteList); i++ {
		fmt.Println(<-c)
	}

}

func checker(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		c <- fmt.Sprintf("%v might be down", link)
		return
	}

	c <- fmt.Sprintf("%v working fine", link)
}
