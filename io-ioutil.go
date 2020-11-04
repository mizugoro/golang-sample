package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	/*
		ioutil.ReadAll()は引数にio.Readerを渡すとその中身をすべて読み出し
		[]byte型で返す。os.Fileの読み出しでは直接Read()を呼んでいたため
		[]byteを用意する必要があったが、これを用いると不要
		file, _ := os.Open("./file.txt")
		message, err := ioutil.ReadAll(file)
	*/

	hello := []byte("hello world\n")
	err := ioutil.WriteFile("./file.txt", hello, 0666)
	if err != nil {
		log.Fatal(err)
	}
	//ファイルの中身をすべて読み出す
	message, err := ioutil.ReadFile("./file.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(message))

}
