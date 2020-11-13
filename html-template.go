package main

import (
	"html/template"
	"net/http"
)

/*ParseFiles(*は戻り値としてエラーを一緒に返すが
Mustを一緒に用いるとエラー時に戻り値ではなくパニックを発生する
一度コンパイルが通ることを確認したテンプレートであれば
毎回エラー処理をする必要性は低いためMust()を利用することがよくある*/
var t = template.Must(template.ParseFiles("index.html"))

func PersonHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close() //処理の最後にBodyを閉じる

	if r.Method == "POST"{
		//リクエストボディをJSONに変換
		var person Person
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&person)

		if err != nil{
			log.Fatal(err)
		}

		//ファイル名を{id}.txtとする
		filename := fmt.Sprintf("%d.txt",person.ID)
		file,err := os.Create(filename) //ファイルを生成
		if err != nil{
			log.Fatal(err)
		}
		defer file.Close()

		//fileにNameを書き込む
		_,err = file.WriteString(person.Name)
		if err != nil {
			log.Fatal(err)
		}

		//レスポンスとしてステータスコード201を送信
		w.WriteHeader(http.StatusCreated)

	}else if r.Method == "GET"{
		//パラメータを取得
		if err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			log.Fatal(err)
		}
		filename := fmt.Sprintf("%d.txt",id)
		b,err := ioutil.ReadFile(filename)
		if err != nil {
			log.Fatal(err)
		}

		//personを生成
		person := Person{
			ID:id,
			Name:String(b),
		}
		t.Execute(w,person)
	}
}
