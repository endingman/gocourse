// main 程序入口点
package main

//import 引入包，外部库整合到你的程序
// go 编译器引用导入包的必要部分，影响程序大小（底层 tree shaking --树瑶） 可执行文件或者最终包里消除未使用或者无效的代码
import "fmt"

//runtime.main_main·f: function main is undeclared in the main package
// func root() {
// 	fmt.Println("hello world")
// }

// package main 必须有一个main函数
func main() {
	fmt.Println("hello world")
}

//编译器将人类语言翻译成计算机可识别的语言，也就是机器码。
// Go Runtime 负责程序的执行，处理诸如内存、分配垃圾回收等，同时协调协程调度
// go 自动处理垃圾回收，不需要手动进行垃圾回收
// go 标准库 构建以及运行所需的基本功能， 提供预先编写的可重用的代码
