package main

import "fmt"

type Task struct {
	ID     int
	Detail string
	done   bool
}

func Finish(task *Task) {
	task.done = true
}
func NewTask(id int, detail string) *Task {
	task := &Task{
		ID:     id,
		Detail: detail,
		done:   false,
	}
	return task
}

func main() {
	//var task Task = Task{1, "buy the milk", true} //Task型
	var task *Task = new(Task) //Taskのポインタ型

	//構造体型もアドレスを取得し、ポインタ型で扱うことができる。変数には構造体の値ではなくアドレスが格納される
	//関数に対して構造体を値渡しするとデータはコピーされるため変更は反映されない
	Finish(task)
	fmt.Println(task.done) //true

	fmt.Println(task.ID)     //1
	fmt.Println(task.Detail) //buy the milk

	//構造体の生成時に値を明示的に指定しなかった場合はゼロ値で初期化されます

	//new()を用いて初期化することができる
	//new()は構造体のフィールドをすべてゼロ値で初期化し、そのポインタを返す
	/*
		ゼロ値
		数値型(int,floatなど):0
		bool型:false
		string型:""空文字列
	*/

	task1 := NewTask(1, "buy the milk")
	// &{ID:1,Detail:buy the milk,done:false}
	fmt.Printf("%+v", task1)

	//コンストラクタにあたる構文がない
	//Newではじまる関数を定義し、その内部で構造体を生成するのが通例
}
