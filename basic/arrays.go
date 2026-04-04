package main

import "fmt"

func main() {

	// var arrayName [size]elementType 数组声明定义
	// size 必须是一个常量表达式，不能是变量或者运行时计算的值
	// elementType 可以是任何有效的 Go 数据类型，包括基本类型、结构体、接口等

	// 声明一个长度为 5 的整数数组
	// 数组是一个固定长度的序列，长度在声明时确定，不能动态改变
	// 数组元素类型必须一致，所有元素都必须是相同的数据类型
	// 数组在内存中是连续存储的，这使得访问数组元素非常高效 ,适合有序数据存储
	// 数组的长度是类型的一部分，因此 [5]int 和 [10]int 是不同的类型
	// 数组是值类型，赋值或者传递数组会创建一个新的副本，而不是引用同一个底层数据结构

	var numbers [5]int //int 类型数组 未定义初始化，默认数组元素值插入为 0

	// 声明并初始化一个字符串数组
	var names = [3]string{"Alice", "Bob", "Charlie"}

	// 访问数组元素
	firstNumber := numbers[0] // 获取第一个元素
	secondName := names[1]    // 获取第二个元素
	fmt.Println("First number:", firstNumber)
	fmt.Println("Second name:", secondName)

	// 修改数组元素
	numbers[2] = 10   // 将第三个元素设置为 10
	names[0] = "Dave" // 将第一个元素修改为 "Dave"

	// 数组长度
	length := len(numbers) // 获取数组的长度
	fmt.Println("Length of numbers array:", length)

	// 多维数组
	var matrix [3][3]int // 声明一个 3x3 的二维整数数组
	matrix[0][0] = 1     // 设置第一个元素为 1
	matrix[1][1] = 5     // 设置第二行第二列的元素为 5
	fmt.Println("Matrix:", matrix)

	originalArray := [3]int{1, 2, 3}
	copiedArray := originalArray // 复制数组，创建一个新的数组
	copiedArray[0] = 10          // 修改复制的数组，不会影响原始数组

	fmt.Println("Original array:", originalArray)
	fmt.Println("Copied array:", copiedArray)

	// array pointer
	arrayPtr := &originalArray // 获取原始数组的指针
	(*arrayPtr)[0] = 20        // 修改通过指针访问的数组元素

	fmt.Println("Original array after modification through pointer:", originalArray)

	//  len 函数获取数组长度 go里面数组长度是固定的，不能动态改变
	for i := 0; i < len(numbers); i++ {
		fmt.Printf("Element at index %d: %d\n", i, numbers[i])
	}

	// for range 循环迭代数组
	for index, value := range names {
		fmt.Printf("Index: %d, Value: %s\n", index, value)
	}

	// for _ range 循环迭代数组，忽略索引
	//_ 是一个特殊的标识符，表示我们不关心这个值
	for _, value := range names {
		fmt.Printf("Value: %s\n", value)
	}

	a, _ := someFunction() // 只获取函数返回的第一个值，忽略第二个值
	fmt.Println("First value from someFunction:", a)

	//  compare arrays
	array1 := [3]int{1, 2, 3}
	array2 := [3]int{1, 2, 3}
	array3 := [3]int{4, 5, 6}

	fmt.Println("array1 == array2:", array1 == array2)
	fmt.Println("array1 == array3:", array1 == array3)
}

func someFunction() (int, int) {
	return 1, 2
}
