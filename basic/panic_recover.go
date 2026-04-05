package main

import "fmt"

func main() {
	//recover 函数用于从 panic 中恢复，允许我们在发生 panic 时捕获错误并继续执行程序。它只能在 defer 函数中调用，并且只能捕获当前 goroutine 中的 panic。

	//当一个函数发生 panic 时，程序会立即停止执行当前函数，并开始逐层向上回退调用栈，直到找到一个 defer 函数中调用了 recover 的地方。如果找到了这样的 defer 函数，recover 会捕获到 panic 的值，并返回该值。此时程序可以继续执行，而不会崩溃。

	//如果没有找到调用 recover 的 defer 函数，或者 recover 没有被正确调用，那么 panic 将继续传播，最终导致程序崩溃。

	//使用 recover 可以帮助我们处理一些不可预见的错误情况，避免程序崩溃，并且可以在发生 panic 时进行一些清理工作，例如关闭文件、释放资源等。

	//需要注意的是，recover 只能捕获当前 goroutine 中的 panic，如果在其他 goroutine 中发生 panic，recover 是无法捕获到的。因此，在使用 recover 时，需要确保它被正确地放置在 defer 函数中，并且只能捕获当前 goroutine 中的 panic。

	//总之，recover 是 Go 语言中用于从 panic 中恢复的函数，可以帮助我们处理错误情况，避免程序崩溃，并且可以在发生 panic 时进行一些清理工作。

	//recover 是返回一个返回值的函数，只有当前对应的协程panic了，才会返回对应的值，恢复成功后，程序继续执行

	process()
	fmt.Println("Program continues after recovering from panic.")
}

func process() {
	// 当前协程程序恢复前执行
	defer func() {
		if r := recover(); r != nil {
			println("Recovered from panic:", r)
		}
	}()
	fmt.Println("start Processing...")
	panic("Something went wrong!")
	// 同样panic后续的代码不会被执行
	fmt.Println("end Processing...")
}

// use cases
//  graceful recovery
//  cleanup
// logging and Reporting

// best practices
// use with defer
// avoid silent Recovery
// avoid overuse
