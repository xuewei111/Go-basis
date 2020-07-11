package main

import "fmt"

func main() {

	//重点讲解 / %
	//说明,如果运算的数都是整数,那么除后,去掉小数,保留整数部分
	fmt.Println(10/4) //2

	var n1 float32 = 10/4 //2
	fmt.Println(n1)

	//如果我们希望保留小数部分,则需要有浮点数参与运算
	var n2 float32 = 10.0/4
	fmt.Println(n2)

	//演示 % 的使用
	fmt.Println("10%3:",10%3)
	fmt.Println("-10%3",-10%3)

	//++ 和 -- 的使用
	var i int = 10
	i++
	fmt.Println(i)

	i--
	fmt.Println(i)
}
