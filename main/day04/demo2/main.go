package main

import "fmt"

//演示golang中string类型的使用
func main() {
	//string的基本使用
	var address string = "南京"

	fmt.Println("address:", address)
	//字符串一旦赋值了,字符串就不能改变了;在Go中字符串是不可变的
	str := "hello"
	//这里就不能去修改str的内容,既go中的字符串不可变
	//str[0] = 'a'
	fmt.Println("str:", str)

	//字符串的两种表示形式
	//1.双引号,会识别转义字符
	//2.反引号:以字符串的原生形式输出,包括换行和特殊字符,可以实现防止攻击,输出源代码等效果

	str2 := "abc\nabc"
	fmt.Println("str2:", str2)

	str3 := `func main() {

	b := false
	fmt.Println("b: ",b)

	//注意事项
	//1.bool类型占用存储空间是一个字节
	fmt.Println("b的占用空间是=",unsafe.Sizeof(b))
	//2.bool类型只能取true或者false
	//var c bool = 1

}`

	fmt.Println("str3:", str3)

	//字符串的拼接方式

	var str4 = "hello" + " world"
	str4 += " hahal"

	fmt.Println("str4:", str4)

	//当一个拼接的操作很长的时候,怎么办？可以分行写
	var str5 = "hello" +
		"hello" +
		" world!"

	fmt.Println("str5:", str5)

	var a int           //0
	var b float32       //0
	var c float64       //0
	var isMarrtied bool //false
	var name string     //""

	//这里的%v 表示按照变量的值输出
	fmt.Printf("a=%d,b=%v,c=%v,isMarrtied=%v,name=%v", a, b, c, isMarrtied, name)
}