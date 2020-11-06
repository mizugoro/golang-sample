package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	urls := []string{
		"http://google.com",
		"http://amazon.com",
		"http://yahoo.com",
	}
	//stringを扱うチャネルを作成
	statusChan := make(chan string)
	for _, url := range urls {
		//取得処理をゴルーチンで行う
		go func(url string) {
			res, err := http.Get(url)
			if err != nil {
				log.Fatal(err)
			}
			defer res.Body.Close()
			//チャネルにstringを書き込む
			statusChan <- res.Status

		}(url)

	}
	for i := 0; i < len(urls); i++ {
		//チャネルからstringを読み出す
		fmt.Println(<-statusChan)
	}
}

/*複数のゴルーチン間でデータをやり取りしたい場合、
channelを用いることで情報をメッセージとして送受信することによって
データを送受信できる*/
