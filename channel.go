package main

import (
	"fmt"
	"log"
	"net/http"
)

func getStatus(urls []string) <-chan string {
	//関数でチャネルを作成
	statusChan := make(chan string)
	for _, url := range urls {
		go func(url string) {
			res, err := http.Get(url)
			if err != nil {
				log.Fatal(err)
			}
			defer res.Body.Close()
			statusChan <- res.Status
		}(url)
	}
	return statusChan
}

func main() {
	urls := []string{
		"http://google.com",
		"http://amazon.com",
		"http://yahoo.com",
	}
	statusChan := getStatus(urls)
	for i := 0; i < len(urls); i++ {
		fmt.Println(<-statusChan)
	}
}

/*複数のゴルーチン間でデータをやり取りしたい場合、
channelを用いることで情報をメッセージとして送受信することによって
データを送受信できる*/
