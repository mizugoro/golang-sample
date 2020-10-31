package main

import "fmt"

func sum(i, j int) {
	//引数がある場合は変数と型を指定します。複数の同じ型が続く場合は型の宣言は一つにまとめることができます
	fmt.Println(i + j)

}

//Goにおいて、関数は複数の値を返すことができる。
//戻り値が複数の場合は型をカンマで区切って指定し丸括弧でくくります
func swap(i, j int) (int, int) {
	return j, i
}
func main() {
	sum(2, 3) //5
	x, y := 3, 4
	x, y = swap(x, y)
	fmt.Println(x, y)
}
