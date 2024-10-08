package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string {
		"http://google.com",
		"http://facebook.com",
		"http://x.com",
		"http://golang.org",
	};

	c := make(chan string);

	for _, link := range links {
		go checkLink(link, c);
	}

	for l := range c{
		go func (link string)  {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
	
}


func checkLink(link string, c chan string) {

	_, err := http.Get(link);

	if err != nil {
		fmt.Println(link, "is not available");
		c <- link;
		return;
	}

	c <- link
	fmt.Println(link, "is available");
}