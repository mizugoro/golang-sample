package main

import "fmt"

func fn(arr [4]string) {
	fmt.Println(arr)
}

func main() {
	var arr1 [4]string
	arr1[0] = "a"
	arr1[1] = "b"
	arr1[2] = "c"
	arr1[3] = "d"
	fmt.Println(arr1[0]) //a

	arr2 := [4]string{"a", "b", "c", "d"}
	arr3 := [...]string{"a", "b", "c", "d"}
	fmt.Println(arr2[1], arr3[2])

	var arr4 [5]string
	fn(arr1)
	//fnt(arr4) //error

}
