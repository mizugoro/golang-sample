package main

import (
	"fmt"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello world")
}

func main() {
	http.HandleFunc("/", IndexHandler) //ルーティングの設定
	/*func HandleFunc(pattern string,
		handler func(ResponsWriter, *Request))

	第一引数はパスのパターンでリクエストを受けたときはこのパターンに一致したハンドラが実行される
	ルートパスに対するリクエストに処理を登録している

	第二引数は二つの引数を受け取る関数になっており、ここではIndexHandlerで実装している
	Requestにはリクエストの情報が入っており、それをもとに組み立てた結果をResponseWriterに書き込むことで
	レスポンスを返せる。io.Writerなので、ここではfmt.Fprint()を用いて文字列を書き込む*/
	http.ListenAndServe(":3030", nil)
	/*ポートを指定してサーバを起動。第二引数は今回使わないためnilを指定している*/
}
