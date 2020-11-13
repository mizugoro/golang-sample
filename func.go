package main

import (
	"errors"
	"fmt"
	"log"
)

/*
名前を付けた戻り値の変数を使うとreturnステートメントに何も書かずに戻すことができる
これを"naked" return と呼ぶ
naked retrunステートメントは短い関数でのみ利用するべき。
*/

func sum(i, j int) {
	//引数がある場合は変数と型を指定します。複数の同じ型が続く場合は型の宣言は一つにまとめることができます
	fmt.Println(i + j)

}

//Goにおいて、関数は複数の値を返すことができる。
//戻り値が複数の場合は型をカンマで区切って指定し丸括弧でくくります
func swap(i, j int) (int, int) {
	return j, i
}

//自作のエラーはerrorsパッケージを用いて作ることができる
func div(i, j int) (int, error) {
	if j == 0 {
		//自作のエラーを返す
		return 0, errors.New("divided by zero")
	}
	return i / j, nil
}

func main() {
	sum(2, 3) //5
	x, y := 3, 4
	x, y = swap(x, y)
	fmt.Println(x, y)

	x, _ = swap(x, y) //第二戻り値を無視する

	//関数が多値を返せることを利用して内部で発生したエラーを戻り値で表現します
	//処理に成功した場合はエラーはnilにし、異常があった場合はエラーだけに値が入り他方はゼロ値になります
	/*
		, err := os.Open("print.go")
		if err != nil {
			//エラー処理

		}
		//fileを用いた処理
	*/
	//関数リテラルを用いると無名関数を作ることができる
	func(i, j int) {
		fmt.Println(i + j)
	}(2, 4)

	n, err := div(10, 0)
	if err != nil {
		//エラーを出力しプログラムを終了する
		log.Fatal(err)
	}
	fmt.Println(n)

}
