package main

// 这里是作业代码_等级3
import "fmt"

func fac(n int) int {
	if n > 1 {
		return n * fac(n-1)
	} else {
		return 1
	}
}

func main() {
	a := 3
	var res3 = fac(a)
	fmt.Println(res3)
	b := 5
	var res5 = fac(b)
	fmt.Println(res5)
}
