package main

import "fmt"

//将一个计算的功能,放到一个函数中,然后在需要使用中,调用即可
func cel(n1 float64,n2 float64,operator byte) float64  {
	var res float64
	switch operator {
	case '+':
		res = n1+n2
	case '-':
		res = n1-n2
	case '*':
		res = n1*n2
	case '/':
		res =n1/n2
	default:
		fmt.Println("错误！")

	}
	return res
}

func main() {



	//需求
	//输入两个数,再输入一个运算符(+,-,*,/),得到结果
	var n1 float64 =1.2
	var n2 float64 =2.6
	var operator byte ='*'

	fmt.Println("res=",cel(n1,n2,operator))

	

}
