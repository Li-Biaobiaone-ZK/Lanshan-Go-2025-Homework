package main

import "fmt"

// 用于记录课上跟着写的代码，作业代码在after class code文件夹内
func main() {

	//practice1
	//你好世界
	fmt.Println("hello world")

	//practice2
	//变量定义
	//手动标注类型
	var a int
	a = 100
	var b int
	b = 200
	fmt.Println(a)
	fmt.Println(b)
	//自动识别类型
	var c = 100
	fmt.Println(c)
	var d = 200
	fmt.Println(d)
	//简写类型
	e := 100
	fmt.Println(e)
	f := 200
	fmt.Println(f)

	//practice3
	//常量定义
	//手动标注类型
	const g int = 100
	fmt.Println(g)
	//自动识别类型
	const pi = 3.1415926
	fmt.Println(pi * 2)

	//practice4
	//for关键字下的循环
	//for 单次表达式：条件表达式：末尾循环体
	for a := 0; a < 10; a++ {
		fmt.Println(a)
	}
	//break
	for b := 0; b < 10; b++ {
		for {
			break
		}
		fmt.Println(b)
	}

}
