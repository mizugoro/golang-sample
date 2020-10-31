package main

import "fmt"

func fn(arr [4]string) {
	fmt.Println(arr)
}

//可変長引数
//関数において引数を次のように指定すると可変長引数として任意の数の引数をその型のスライスとして受け取ることができる

func sum(nums ...int) (result int) {
	//numsは[]int型
	for _, n := range nums {
		result += n
	}
	return
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

	fmt.Println(sum(1, 2, 3, 4)) //10

}
