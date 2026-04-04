package main

import "fmt"

func main() {
	// if condition {
	// 	block of code
	// }

	// 无论满足哪个条件都会返回 x 或者 y 时，
	// 	if condition {
	//     return x
	// }
	// return y

	age := 25
	if age > 18 {
		fmt.Println("You are an adult")
	}

	// 	if condition {
	// do something
	// } else {
	// do something
	// }

	// if condition1 {
	// 	block of code1
	// } else if condition2 {
	// 	block of code2
	// }else {
	// 	block of code3
	// }

	// || OR 满足一个就返回
	//  && AND 满足都返回
	// ! NOT
}
