package main

import (
	"errors"
	"fmt"
)

//エラーは戻り値によって表現するのが基本ですが、そうではない場合もある
//配列やスライスの範囲外にアクセスした場合
//代わりにパニックという方法でエラーが発生します

//このパニックで発生したエラーはrecoverという組み込み関数で取得しそこでエラーを実施する
//recoverをdeferの中に書くことでパニックで発生したエラーの処理を実施してから関数を抜けることができる

func main() {
	/*
		defer func() {
			err := recover()
			if err != nil {
				//runtime error:index out of range
				log.Fatal(err)
			}
		}()

		a := []int{1, 2, 3}
		fmt.Println(a[10]) //パニックが発生
	*/
	a := []int{1, 2, 3}
	for i := 0; i < 10; i++ {
		if i >= len(a) {
			panic(errors.New("index out of range"))
		}
		fmt.Println(a[i])
	}
}
