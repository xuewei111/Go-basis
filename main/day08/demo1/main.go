package main

import "fmt"

func main() {


	var age int8
	fmt.Println("请输入你的年龄！")
	fmt.Scanln(&age)

	if age<18{
		fmt.Println("玩")
	}else {
		fmt.Println("da1")
	}
}
