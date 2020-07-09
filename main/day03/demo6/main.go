package main

import (
	"fmt"
	"unsafe"
)

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

	//整型的使用细节
	var n1 = 100 //? n1 是什么类型

	//这里我们给介绍一下如何查看某个变量的数据类型
	//fmt.Parintf() 可以用于做格式化输出
	fmt.Printf("n1 的类型是%T \n",n1)

	//如何在程序中查看某个变量的占用字节大小和数据类型(使用较多)
	var n2 int64 = 10
	fmt.Printf("n2 的类型 %T n2占用的字节数是 %d \n",n2,unsafe.Sizeof(n2))


	//golang程序中整型变量在使用时,遵守保小不保打的原则
	//既:在保证程序正确运行下,尽量使用占用空间小的数据类型
	var age byte = 12
	fmt.Println("age=",age);
}
