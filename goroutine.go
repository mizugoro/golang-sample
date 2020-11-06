package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

func main() {
	wait := new(sync.WaitGroup)
	urls := []string{
		"http://github.com",
		"http://google.com",
		"http://amazon.com",
	}
	for _, url := range urls {
		//取得処理をゴルーチンで実行する
		//関数の前にgoというキーワードを加えるとそれだけで
		//関数の実行が別のゴルーチンで行われるようになる
		wait.Add(1)
		go func(url string) {
			res, err := http.Get(url)
			if err != nil {
				log.Fatal(err)
			}
			defer res.Body.Close()
			fmt.Println(url, res.Status)
			wait.Done()
		}(url)
	}
	//待ち合わせ
	wait.Wait()
	//main()が終わらないように待ち合わせる
	//time.Sleep(time.Second)
}

//起動したゴルーチンの終了を待ち合わせるにはsync.WaitGroup()が利用できる
//sync.WaitGroupはAdd()でカウントを増やしDone()でカウントを減らしWait()でカウントが
//ゼロになるまで待ち合わせる
