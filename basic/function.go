package main

import "fmt"

func main() {
	// func functionName(parameters) (returnTypes) {
	// 	block of code
	// }
	// 外部函数需要使用大写字母开头，才能被其他包访问到（导出）。如果函数名以小写字母开头，则只能在同一包内访问（未导出）。

	fmt.Println(add(10, 20))

	// 匿名函数 又叫闭包 没有名字，可以直接定义和调用，或者赋值给一个变量以便后续调用。
	func() {
		fmt.Println("This is an anonymous function")
	}() // 定义并立即调用匿名函数

	result := func(a int, b int) int {
		return a * b
	}(5, 10) // 定义并立即调用匿名函数，传入参数 5 和 10
	fmt.Println(result) // 输出: 50

	// 函数也是一种类型，可以作为参数传递给其他函数，或者作为返回值返回。
	// 例如，下面的函数接受一个函数作为参数，并调用它：
	execute := func(f func(int, int) int, x int, y int) int {
		return f(x, y)
	}
	sum := execute(add, 10, 20) // 将 add 函数作为参数传递给 execute 函数
	fmt.Println(sum)            // 输出: 30

	operation := add
	add_result := operation(15, 25) // 将 add 函数赋值给变量 operation，并调用它
	fmt.Println(add_result)         // 输出: 40

	// 函数作为参数传递 passing a function as an argument
	result = applyOperation(10, 20, add) // 将 add 函数作为参数传递给 applyOperation 函数
	fmt.Println(result)                  // 输出: 30

	// 函数作为返回值返回 returning and using a function
	multiplier := createMultiplier(3) // 创建一个乘法器函数，乘数为 3
	fmt.Println(multiplier(10))       // 输出: 30，调用乘法器函数，传入参数 10
}

func add(a int, b int) int {
	return a + b
}

// 函数也是一种类型，可以作为参数传递给其他函数，或者作为返回值返回。
// 例如，下面的函数接受一个函数作为参数，并调用它：
// function that takes another function as a parameter
func applyOperation(a int, b int, operation func(int, int) int) int {
	return operation(a, b)
}

// function that returns a function
func createMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}
