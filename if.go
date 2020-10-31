package main

import (
	"fmt"
)

func main() {
	a, b := 10, 100
	if a > b {
		fmt.Println("a is larger than b")
	} else if a < b {
		fmt.Println("a is smaller than b")
	} else {
		fmt.Println("a equals b")
	}
	//if文の条件部に丸括弧は必要ない

	//if,else文が繰り返す場合はswitch文を用いたほうがすっきりする場合がある
	//値の比較だけでなく、条件分岐にも使用できる

	n := 10
	switch n {
	case 15:
		fmt.Println("fizzbuzz")

	case 5, 10:
		fmt.Println("buzz")

	case 3, 6, 9:
		fmt.Println("fizz")

	default:
		fmt.Println(n)
	}
	//Cなどでは1つのcaseが実行されるとその次のcaseに処理が移るため、breakを書く必要があった
	//Goではcaseが１つ実行されると次のcaseに移らないためbreakを書く必要がない

	//ただcaseが終わった後に次のcaseに移りたい場合はfallthroughを書くことで明示的に
	//つぎのcaseに処理を進めることができる
	n = 3
	switch n {
	case 3:
		n = n - 1
		fallthrough
	case 2:
		n = n - 1
		fallthrough
	case 1:
		n = n - 1
		fmt.Println(n)

	}
	//caseに値だけでなく式も指定できます
	c := 10
	switch {
	case c%15 == 0:
		fmt.Println("FizzBuzz")
	case c%5 == 0:
		println("Buzz")
	case c%3 == 0:
		println("Fizz")
	default:
		println(c)

	}
}
