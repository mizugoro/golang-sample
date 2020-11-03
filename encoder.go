package main

import (
	"encoding/json"
	"log"
	"os"
)

type Person struct {
	ID      int    `json:"id"` //idというキーで格納する
	Name    string `json:"name"`
	Email   string `json:"-"` //JSONに格納しない
	Age     int    `json:"age"`
	Address string `json:"address,omitempty"` //値が空なら無視する
	memo    string
}

func main() {
	person := &Person{
		ID:      1,
		Name:    "Gopher",
		Email:   "gopher@example.org",
		Age:     5,
		Address: "",
		memo:    "golang lover",
	}

	//ファイルを開く
	file, err := os.Create("./person.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//エンコーダーの取得
	encoder := json.NewEncoder(file)

	//JSONエンコードしたデータの書き込み
	err = encoder.Encode(person)
	if err != nil {
		log.Fatal(err)
	}

}
