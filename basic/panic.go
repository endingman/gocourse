package main

import "fmt"

func main() {
	//panic 是 Go 语言中的一个内置函数，用于引发运行时错误。当程序遇到无法继续执行的情况时，可以调用 panic 来中断程序的正常流程，并输出一个错误信息。panic 通常用于处理不可恢复的错误，例如数组越界、空指针引用等。

	//当 panic 被调用时，程序会立即停止执行当前函数，并开始逐层向上回退调用栈，直到找到一个 defer 语句或者 main 函数的结束位置。如果在回退过程中遇到了 defer 语句，那么 defer 语句中的函数会被执行，然后继续回退，直到找到一个能够处理 panic 的 defer 语句或者 main 函数的结束位置。

	//如果在回退过程中没有找到任何能够处理 panic 的 defer 语句，那么程序将会崩溃，并输出 panic 的错误信息和调用栈信息。

	//panic 的语法格式如下：
	// panic(expression)

	//在实际开发中，我们应该尽量避免使用 panic，因为它会导致程序崩溃，影响用户体验。相反，我们应该使用错误处理机制来处理可预见的错误情况，例如返回错误值或者使用 error 类型来表示错误。

	//当然，在某些情况下，panic 可能是不可避免的，例如在处理一些严重的错误或者在开发过程中进行调试时。但是，我们应该谨慎使用 panic，并确保在适当的地方使用 defer 来处理 panic，以避免程序崩溃和资源泄漏。

	process(10)
	process(-3)
}

func process(input int) {
	defer fmt.Println("Deferred 1")
	defer fmt.Println("Deferred 2")

	if input < 0 {
		// 在panic之后的代码都不会被执行，但是在panic之前的defer语句会被执行
		panic("before panic")
	}
	fmt.Println("Processing input:", input)

	defer fmt.Println("Deferred 3")

}
