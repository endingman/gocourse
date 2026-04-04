// 运算符
package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b int = 10, 3
	var result int

	result = a + b
	fmt.Println("Addition:", result)

	result = a - b
	fmt.Println("Subtraction:", result)

	result = a * b
	fmt.Println("Multiplication:", result)

	result = a / b
	fmt.Println("Division:", result)

	result = a % b
	fmt.Println("Remainder:", result)

	// float 除法结果， 必须其一是浮点数
	const p float64 = 22 / 7.0
	fmt.Println("Pi:", p)

	//  overflow with unsigned integers
	var maxInt int = 1<<63 - 1
	fmt.Println("Max int:", maxInt)
	maxInt = maxInt + 1
	fmt.Println("Max int:", maxInt)

	// max value for unit64 type
	var uintMax uint = 1<<64 - 1
	fmt.Println("Max uint:", uintMax)
	// 上溢出 0
	uintMax = uintMax + 1
	fmt.Println("Max uint:", uintMax)

	// underflow with floats number
	var smallFloat float64 = 1.0e-323
	fmt.Println("Small float:", smallFloat)
	// 下溢出 0
	smallFloat = smallFloat / math.MaxFloat64
	fmt.Println("Small float:", smallFloat)

}
