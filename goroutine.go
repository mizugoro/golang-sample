package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	urls := []string{
		"http://github.com",
		"http://google.com",
		"http://amazon.com",
	}
	for _, url := range urls {
		res, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()
		fmt.Println(url, res.Status)
	}
}
