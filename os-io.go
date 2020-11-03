package main

import (
	"fmt"
	"log"
	"os"
)

type Writer interface {
	Write(p []byte) (n int, err error)
}

//ファイルの読み出し
type Reader interface {
	Read(p []byte) (n int, err error)
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

	//ファイルを開く
	file, err = os.Open("./file.txt")
	if err != nil {
		log.Fatal(err)
	}
	//プログラムが終わったらファイルを閉じる
	defer file.Close()

	message1 := make([]byte, 12)

	//ファイル内のデータをスライスに読みだす
	_, err = file.Read(message1)
	if err != nil {
		log.Fatal(err)
	}
	//文字列にして表示
	fmt.Println(string(message1))

}
