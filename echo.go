package main

import (
	"flag"
	"os"
)

var omitNewline = flag.Bool("n", false, "don't print final newline")

const (
	Space   = " "  //空白
	Newline = "\n" //改行
)

func main() {
	flag.Parse() //パラメータリストを調べてflagに設定
	var s string = ""
	for i := 0; i < flag.NArg(); i++ {
		if i > 0 {
			s += Space //空白を追加
		}
		s += flag.Arg(i)
	}
	if !*omitNewline {
		s += Newline //改行を追加
	}
	os.Stdout.WriteString(s)
}
