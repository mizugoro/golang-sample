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

func main() {
	//var task Task = Task{1, "buy the milk", true} //Task型
	var task *Task = &Task{done: false} //Taskのポインタ型
	//構造体型もアドレスを取得し、ポインタ型で扱うことができる。変数には構造体の値ではなくアドレスが格納される
	//関数に対して構造体を値渡しするとデータはコピーされるため変更は反映されない
	Finish(task)
	fmt.Println(task.done) //true
	/*
		fmt.Println(task.ID)     //1
		fmt.Println(task.Detail) //buy the milk

	*/

	//構造体の生成時に値を明示的に指定しなかった場合はゼロ値で初期化されます
}
