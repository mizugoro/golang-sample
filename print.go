package main

/*Goのコードはパッケージの宣言から始まる
プログラムをコンパイルして実行するとまずsampleパッケージの中にあるmain関数が実行されるため、
ここに処理を記述している。
*/

import (
	"fmt"
)

/*importはプログラム内にほかのパッケージを取り込むために使用します
improtしたパッケージはドットを付けてアクセスできます。
オプションの指定
f "fmt"とした場合はfがfmtを表すようになる
_ "github.com/~~~"
_を付けた場合はimportしたpackageを使用しないことをコンパイラに伝える
.を付けた場合は使用時にパッケージ名が省略可能になります*/

func main() {
	// var message string = "hello world"
	/*var 変数名 型で変数の宣言を行う
	ex: var foo,bar, buz string = "foo","bar","buz
	1度に多くの変数を同じ型で宣言する場合はvarと二つ目以降の型を省略することができる
	*/

	message := "hello world"

	//この書き方をした場合、変数の型は代入する値からコンパイラによって推論される。
	fmt.Println(message)

	//変数を宣言し明示的に値を初期化しなかった場合は変数はゼロ値というデフォルト値で初期化される
	//intのゼロ値は0であるため次のコードは０を出力する
	var i int
	fmt.Println(i)
}
