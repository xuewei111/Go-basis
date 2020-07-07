package main

import "fmt"

//演示Golang中整数类型的使用
func main()  {

	i := 1
	fmt.Println("i=",i)
	//测试一下int8的范围(127~-128)
	//var j int8 =128
	//fmt.Println("j=",j)

	//int的无符号
	//测试一下uint8的范围(0-255),其他的uint16,uint32,uint64一样
	//overflows
	var k uint16 = 256
	fmt.Println("k=",k)

	//int,uint,rune,byte的使用
	var a int = 8900
	fmt.Println("a:",a)

	//var b uint = -1
	//fmt.Println("b:",b)
	//var c byte = "a"
	//fmt.Println("c:",c)
}
