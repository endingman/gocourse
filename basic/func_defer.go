package main

import "fmt"

func main() {
	// 关键字 defer 允许我们推迟到函数返回之前（或任意位置执行 return 语句之后）一刻才执行某个语句或函数（为什么要在返回之后才执行这些语句？因为 return 语句同样可以包含一些操作，而不是单纯地返回某个值）。

	// defer 和 return 的执行顺序是先为返回值赋值，然后执行 defer，然后 return 到函数调用处。)

	// 关键字 defer 的用法类似于面向对象编程语言 Java 和 C# 的 finally 语句块，它一般用于释放某些已分配的资源。

	// defer 语句会将函数调用推迟到包含它的函数返回之后执行。被 defer 推迟的函数调用会按照先进后出（LIFO）的顺序执行。类似于栈结构，最后被 defer 的函数调用会最先执行。

	// defer 语句常用于确保一些重要的清理操作能够执行，例如关闭文件、释放锁、关闭数据库连接等。无论函数是正常返回还是由于 panic 导致的异常退出，defer 语句都能保证被推迟的函数调用会被执行，从而帮助我们避免资源泄漏和保持程序的健壮性。

	// defer 语句的语法格式如下：
	// defer statement

	process(1)
}

func process(i int) {
	defer fmt.Println("First defer statement executed")
	defer fmt.Println("Second defer statement executed")
	defer fmt.Println("Third defer statement executed")

	// 先返回值，再执行 defer 语句
	i++

	fmt.Println("Normal execution of the function")
	fmt.Println("Value of i:", i)
}
