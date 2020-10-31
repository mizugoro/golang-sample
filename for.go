package main

import "fmt"

func main() {
	//条件部の丸括弧は必要ない
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	//繰り返しを表現する方法はfor文以外になくwhile文といった表現はすべてfor文を用いて表現する

	/*C言語の場合
	intn = 0;
	while(n < 10){
		printf("n = %d\n",n);
		n++;
	}
	*/
	//Go言語の場合
	n := 0
	for n < 10 {
		fmt.Println(n)
		n++
	}

	//無限ループ
	/*
		in C
		for(;;){
			dosomething();
		}

		in Go
		for {
			dosomething()
		}
	*/

	//繰り返し制御はbreak,continueを使う
	x := 0
	for {
		x++
		if x < 10 {
			break
		}
		if x%2 == 0 {
			continue
		}
		fmt.Println(x)
	}
}
