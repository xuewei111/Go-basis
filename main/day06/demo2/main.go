package main

import "fmt"

func main() {

	//假如还有97天放假：问：XX个星期零X天
	var days int = 97

	var week int = days/7
	var day int =  days % 7

	fmt.Printf("%d个星期%d天\n",week,day)

}
