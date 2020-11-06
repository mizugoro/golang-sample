package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
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
	//１秒後に値が読み出せるチャネル
	timeout := time.After(time.Second)
	urls := []string{
		"http://amazon.com",
		"http://google.com",
		"http://github.com",
	}
	statusChan := getStatus(urls)
LOOP: //ラベル名は任意
	for {
		select {
		case status := <-statusChan:
			fmt.Println(status) //受信したデータを表示する
		case <-timeout:
			break LOOP //このfor/selectを抜ける
		}
	}
}
