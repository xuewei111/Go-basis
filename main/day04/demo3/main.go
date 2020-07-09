package main

import "fmt"

//演示golang中基本数据类型的转换
func main() {

	var i int32 = 100
	//希望将I =>folat
	var n1 float32 = float32(i)
	var n2 int8 = int8(i)
	var n3 int64 = int64(i)  //低精度->高精度

	fmt.Println("n1: i: n2: n3: ",n1,i,n2,n3)
	fmt.Printf("i type is %T",i)

	//在转换中,比如讲 int64 转成 int8 (-128~127) ,编译时不会报错
	//只是转换的结果是按溢出处理的,和我们希望的结果表示太一样
	var num1 int64 = 999999
	var num2 int8 = int8(num1)
	fmt.Println("num2: ",num2)

}
