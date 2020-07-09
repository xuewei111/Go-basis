package main

import (
	"fmt"
	"unsafe"
)

//演示golang中bool类型使用
func main() {

	b := false
	fmt.Println("b: ",b)

	//注意事项
	//1.bool类型占用存储空间是一个字节
	fmt.Println("b的占用空间是=",unsafe.Sizeof(b))
	//2.bool类型只能取true或者false
	//var c bool = 1

}
