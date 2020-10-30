package sample

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
	fmt.Println("Hello world.")
}
