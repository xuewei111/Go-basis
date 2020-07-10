package main

import "fmt"

func main() {
	var num int = 9
	fmt.Printf("num address=%v\n",&num)

	var ptr *int
	ptr = &num
	*ptr = 10
	fmt.Printf("ptr的地址是: %v\n",ptr)
	fmt.Printf("ptr指向的值是: %v",*ptr)

}
