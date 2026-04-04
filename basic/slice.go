package main

import (
	"fmt"
	"slices"
)

func main() {
	// var sliceName[] elementType
	// sliceName := []elementType

	// slice 是数组的一个视图（引用，并不存储数据），底层是一个数组，slice 包含了指向数组的指针、长度和容量 切片是一个 长度可变的数组

	// 注意 绝对不要用指针指向 slice。切片本身已经是一个引用类型，所以它本身就是一个指针！！

	// 所以下面两种方法可以生成相同的切片:
	// make([]int, 50, 100)
	// new([100]int)[0:50]

	// slice 可以动态增长，长度和容量可以不同，长度是当前元素的数量，容量是底层数组的大小

	// slice 的零值是 nil，长度和容量都是 0

	// append 函数用于向 slice 添加元素，如果超过容量会自动扩容

	// copy 函数用于复制 slice 的元素到另一个 slice 中

	// range 关键字用于迭代 slice 的元素，返回索引和值

	// make 函数用于创建 slice，指定长度和容量 sliceName := make([]elementType, length, capacity)

	// 切片在内存中的组织方式实际上是一个有 3 个域的结构体：指向相关数组的指针，切片长度以及切片容量-slice.png

	// new () 和 make () 的区别
	// 看起来二者没有什么区别，都在堆上分配内存，但是它们的行为不同，适用于不同的类型。

	// new (T) 为每个新的类型 T 分配一片内存，初始化为 0 并且返回类型为 * T 的内存地址：这种方法 返回一个指向类型为 T，值为 0 的地址的指针，它适用于值类型如数组和结构体（参见第 10 章）；它相当于 &T{}。
	// make(T) 返回一个类型为 T 的初始值，它只适用于 3 种内建的引用类型：切片、map 和 channel
	// 换言之，new 函数分配内存，make 函数初始化；new and make.png

	a := []int{1, 2, 3, 4, 5}     // 声明并初始化一个整数 slice
	slice1 := a[1:4]              // 创建一个 slice，包含 a 中索引 1 到 3 的元素 包含索引 1 不包含索引 4的值 即索引值含左不含右
	fmt.Println("Slice:", slice1) // 输出: Slice: [2 3 4]
	slice1 = append(slice1, 6, 7) // 向 slice 添加元素，可能会触发底层数组的扩容

	// append using and useful
	// 	将切片 b 的元素追加到切片 a 之后：a = append(a, b...)

	// 复制切片 a 的元素到新的切片 b 上：

	//  b = make([]T, len(a))
	//  copy(b, a)
	// 删除位于索引 i 的元素：a = append(a[:i], a[i+1:]...)

	// 切除切片 a 中从索引 i 至 j 位置的元素：a = append(a[:i], a[j:]...)

	// 为切片 a 扩展 j 个元素长度：a = append(a, make([]T, j)...)

	// 在索引 i 的位置插入元素 x：a = append(a[:i], append([]T{x}, a[i:]...)...)

	// 在索引 i 的位置插入长度为 j 的新切片：a = append(a[:i], append(make([]T, j), a[i:]...)...)

	// 在索引 i 的位置插入切片 b 的所有元素：a = append(a[:i], append(b, a[i:]...)...)

	// 取出位于切片 a 最末尾的元素 x：x, a = a[len(a)-1:], a[:len(a)-1]

	// 将元素 x 追加到切片 a：a = append(a, x)

	sliceCopy := make([]int, len(slice1))   // 创建一个新的 slice，长度与 slice1 相同
	copy(sliceCopy, slice1)                 // 复制 slice1 的元素到 sliceCopy 中
	fmt.Println("Copied Slice:", sliceCopy) // 输出: Copied Slice: [2 3 4 6 7]

	if slices.Equal(slice1, sliceCopy) {
		fmt.Println("Slice1 and SliceCopy are equal")
	}

	// nil slice
	var nilSlice []int
	fmt.Println("Nil Slice:", nilSlice)                  // 输出: Nil Slice: []
	fmt.Println("Length of Nil Slice:", len(nilSlice))   // 输出: Length of Nil Slice: 0
	fmt.Println("Capacity of Nil Slice:", cap(nilSlice)) // 输出: Capacity of Nil Slice: 0

	// empty slice
	emptySlice := []int{}
	fmt.Println("Empty Slice:", emptySlice)                  // 输出: Empty Slice: []
	fmt.Println("Length of Empty Slice:", len(emptySlice))   // 输出: Length of Empty Slice: 0
	fmt.Println("Capacity of Empty Slice:", cap(emptySlice)) // 输出: Capacity of Empty Slice: 0

	// slice 的切片操作
	slice2 := a[:3]                // 从开始到索引 2 的元素
	slice3 := a[2:]                // 从索引 2 到结束的元素
	fmt.Println("Slice2:", slice2) // 输出: Slice2: [1 2 3]
	fmt.Println("Slice3:", slice3) // 输出: Slice3: [3 4 5]

	// 在内存中，一个字符串实际上是一个双字结构，即一个指向实际数据的指针和记录字符串长度的整数（见图 string-slice.png）。因为指针对用户来说是完全不可见，因此我们可以依旧把字符串看做是一个值类型，也就是一个字符数组。

}
