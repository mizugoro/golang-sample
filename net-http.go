package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Person struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello world")
}
func PersonHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close() //処理の最後にBodyを閉じる

	if r.Method == "POST" {
		//リクエストボディをJSONに変換
		var person Person
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&person)
		if err != nil {
			log.Fatal(err)
		}
		//ファイル名を{id}.txtとする
		filename := fmt.Sprint("%d.txt", person.ID)
		file, err := os.Create(filename)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		//ファイルにNameを書き込む
		_, err = file.WriteString(person.Name)
		if err != nil {
			log.Fatal(err)
		}
		//
		w.WriteHeader(http.StatusCreated)
	}
}

func main() {
	http.HandleFunc("/", IndexHandler)         //ルーティングの設定
	http.HandleFunc("/persons", PersonHandler) //ルーティングの設定
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
