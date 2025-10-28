package main

// 这里是作业代码_等级X
import "fmt"

func main() {
	var sum int = 0
	var index int = 0
	var num int
	fmt.Scanln(&num)
	for num != 0 {
		index++
		sum += num
		fmt.Scanln(&num)
	}
	var res float64 = float64(sum / index)
	fmt.Println(res)
}
