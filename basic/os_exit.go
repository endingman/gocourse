package main

import (
	"fmt"
	"os"
)

func main() {
	// os.Exit(0)  // 退出程序，参数为退出状态码，0 通常表示成功，非零表示错误
	// 程序退出后，defer 语句不会被执行，所有的资源会被立即释放，程序会直接终止，不会执行任何剩余的代码。
	fmt.Println("Starting the main function.")
	// exit with status code of 0, indicating successful execution
	os.Exit(0)
	// This will not never executed.
	fmt.Println("End of the main function.")
}
