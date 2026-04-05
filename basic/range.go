package main

import "fmt"

func main() {
	message := "Hello, world!"
	// range 关键字操作的是数据结构的副本，而不是原始数据结构本身。在迭代过程中，range 会创建一个新的副本来存储当前迭代的元素，不会对原始数据结构的原始值修改。

	// 如果不打算获取迭代的索引或者值，可以使用 _ 来忽略它。
	for index, char := range message {
		// fmt.Println(index, char)                              // index 是字符在字符串中的位置，char 是字符的 Unicode 码点（rune 类型）
		fmt.Printf("Index: %v, Character: %c\n", index, char) // 使用 %c 格式化字符
	}
}
