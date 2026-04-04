package main

import "fmt"

// 外部定义 ，全部变量
var (
	FirstName = "John"
)

// := 只能在函数内部使用， 即作用域内
// middleName := "Smith"

func main() {
	// var 初始化定义
	var age int //未初始化变量， int 默认值为 0 ，bool 默认值为 false string 默认值为 ""  pointer slices maps functions ans structs 默认值为 nil
	var name string = "John"

	//:= 初始化赋值
	count := 10
	lastName := "Smith"

	fmt.Println(age, name)
	fmt.Println(count, lastName)
	fmt.Println(FirstName)
}
