package main

import (
	"encoding/json"
	"fmt"
	"log"
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
	b, err := json.Marshal(person)
	//ポインタをjson.Marshal()に渡すだけでデフォルトのフォーマットで
	//JSON文字列の[]byteを生成できる
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b)) //文字列に変換

	//JSONから構造体への変換

	var person1 Person
	c := []byte(`{"id":1,"name":"Gopher","age":5}`)
	err = json.Unmarshal(c, &person1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(person1)
}
