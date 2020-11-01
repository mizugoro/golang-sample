//マップは値をkey-valueの対応で保存するデータ構造

package main

import "fmt"

func main() {
	//intのキーにstringの値を格納するMAPは次のように宣言する
	var month map[int]string = map[int]string{}

	//次のようにキーを指定する
	month[1] = "January"
	month[2] = "February"
	fmt.Println(month) //map[1:January,2:February]

	//宣言と初期化を同時に行う場合
	week := map[int]string{
		1: "monday",
		2: "tuesday",
	}
	fmt.Println(week) //map[1:monday,2:tuesday]

	//mapから値を取り出す場合
	jan := month[1]
	fmt.Println(jan) //january

	_, ok := month[1]
	if ok {
		//データがあった場合の処理
		fmt.Println("Ok,Google")
	}

	//mapからデータを消す場合は組み込み関数のdelete()を利用する
	delete(month, 2)
	fmt.Println(month)

}
