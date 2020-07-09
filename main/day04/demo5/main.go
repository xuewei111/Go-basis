package main

import (
	"fmt"
	"strconv"
)

//演示golang中string转成基本数据类型
func main() {
	var str string = "true"
	var b bool

	//b,_ = strconv.ParseBool(str)
	//说明
	//1.strconv.ParseBool(str) 函数会返回两个值(value bool,err error)
	//2.因为我只想获取到value bool,不想获取到err 所以我使用_忽略
	b,_ = strconv.ParseBool(str);

	fmt.Printf("b type %T b=%v\n",b,b)

	var str2 string = "123456789"
	var n1 int64
	n1,_ = strconv.ParseInt(str2,10,64)
	fmt.Printf("n1 type %T n1 = %v\n",n1,n1)

	var str3 string = "123.456"
	var n3 float64
	n3,_ = strconv.ParseFloat(str3,64)
	fmt.Printf("n3 type %T n3 = %v\n",n3,n3)



	//注意
	//如果没有转成功则变成默认值
	var str4 string = "hello"
	n1,_ = strconv.ParseInt(str4,10,64)
	fmt.Printf("n1 type %T n1 = %v\n",n1,n1)

}
