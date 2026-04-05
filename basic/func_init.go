package main

import "fmt"

func main() {
	//init 函数是一个特殊的函数，在 Go 程序中被自动调用，用于在程序启动时进行初始化工作。每个包都可以定义一个或多个 init 函数，这些函数会在 main 函数执行之前被调用。init 函数通常用于设置包级变量、注册类型、执行必要的初始化逻辑等。

	//init 函数的特点：

	//1. 自动调用：init 函数会在程序启动时自动调用，无需显式调用。
	//2. 多个 init 函数：一个包可以定义多个 init 函数，它们会按照在源代码中出现的顺序被调用。
	//3. 无参数和返回值：init 函数不能接受参数，也不能返回值。
	//4. 包级作用域：init 函数只能访问包级变量，不能访问其他包的变量。

	//init 函数的使用场景：

	//1. 初始化包级变量：可以在 init 函数中设置包级变量的初始值。
	//2. 注册类型或函数：可以在 init 函数中注册类型、函数或其他资源，以便在程序运行时使用。
	//3. 执行必要的初始化逻辑：可以在 init 函数中执行一些必要的初始化逻辑，例如连接数据库、加载配置文件等。

	//需要注意的是，init 函数的执行顺序是由 Go 语言运行时决定的，通常情况下，init 函数会按照包的导入顺序被调用。因此，如果一个包依赖于另一个包，那么被依赖的包的 init 函数会先被调用。

	// 总之，init 函数是 Go 语言中一个重要的特性，用于在程序启动时进行必要的初始化工作，确保程序在运行时能够正确地设置和配置所需的资源。

	fmt.Println("Main function is running...")
}

func init() {
	fmt.Println("Initializing Package1...")
}
func init() {
	fmt.Println("Initializing Package2...")
}
func init() {
	fmt.Println("Initializing Package3...")
}
