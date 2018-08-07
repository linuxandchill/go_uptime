package main

import (
	"fmt"
	"net/http"
	"time"
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

	//wait for value to be received on the channel
	// continue pinging
	for {
		go func(link string) {
			time.Sleep(3 * time.Second)
			checker(link, c)
		}(<-c)
	}

	/*
		less readable IMO, not clear that we're waiting for val from channel
		for link := range c{
			go checker(link, c)
		}
	*/

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
