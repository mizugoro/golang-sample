package main

import (
	"log"
	"os"
)

type Writer interface {
	Write(p []byte) (n int, err error)
}

func main() {
	//ファイルの生成
	file, err := os.Create("./file.txt")
	if err != nil {
		log.Fatal(err)
	}
	//プログラムが終わったらファイルを閉じる
	defer file.Close()

	//書き込むデータを[]byteで用意する
	message := []byte("hello world\n")

	//write()を用いて書き込む
	_, err = file.Write(message)
	if err != nil {
		log.Fatal(err)
	}

}
