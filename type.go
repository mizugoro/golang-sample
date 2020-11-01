package main

//IDと優先度を取得してタスクを処理する関数を考える
//typeを用いて既存の型を拡張する,typeの後には型の名前、その型の定義が続く
//
type ID int
type Priority int

func ProcessTask(id, priority int) {

}
func main() {
	var id ID = 3
	var priority Priority = 5
	//呼び出す際には型が適合していないとコンパイルエラーになる
	ProcessTask(priority, id) //コンパイルエラー

	//このように適切な方を用意することで型レベルの整合性を
	//コンパイル時にチェックでき堅牢なプログラムを記述できる
}
