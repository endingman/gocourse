package main

import (
	"fmt"
	"sort"
)

func main() {
	// map 是一种无序的键值对集合,不管是按照 key 还是按照 value 默认都不排序(如果你想为 map 排序，需要将 key（或者 value）拷贝到一个切片，再对切片排序)，提供了快速的查找、插入和删除操作
	// map 是引用类型，需要使用 make 函数创建一个新的 map，或者使用 map 字面量来初始化一个 map
	// var mapVariable map[keyType]valueType
	// mapVariable := make(map[keyType]valueType) 创建一个新的 map

	// 	在声明的时候不需要知道 map 的长度，map 是可以动态增长的。

	// 未初始化的 map 的值是 nil。

	// map 的零值是 nil，不能直接使用，需要使用 make 函数创建一个新的 map

	// map 的键必须是可比较的类型，值可以是任意类型
	// map 的访问和修改使用索引语法，map[key] = value 添加或更新键值对，value = map[key] 访问值，delete(map, key) 删除键值对

	// map 的迭代使用 range 关键字，返回键和值
	// for key, value := range mapVariable {
	// 	// 处理键值对
	// }

	// map 的长度使用 len 函数获取，map 的容量没有固定的限制，会根据需要动态扩展

	// map 是引用类型，多个变量可以引用同一个 map，修改其中一个变量会影响其他变量

	// map 的性能在大多数情况下是非常好的，但在某些情况下可能会出现性能问题，例如大量的哈希冲突或者频繁的删除操作

	// map 的底层实现是哈希表，使用哈希函数将键映射到数组索引，处理哈希冲突的方法是链地址法，即在每个数组索引处维护一个链表来存储具有相同哈希值的键值对

	var (
		barVal = map[string]int{"alpha": 34, "bravo": 56, "charlie": 23,
			"delta": 87, "echo": 56, "foxtrot": 12,
			"golf": 34, "hotel": 16, "indio": 87,
			"juliet": 65, "kiln": 43, "lima": 98}
	)

	fmt.Println("unsorted:")
	for k, v := range barVal {
		fmt.Printf("Key: %v, Value: %v / ", k, v)
	}
	keys := make([]string, len(barVal))
	i := 0
	for k, _ := range barVal {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	fmt.Println()
	fmt.Println("sorted:")
	for _, k := range keys {
		fmt.Printf("Key: %v, Value: %v / ", k, barVal[k])
	}

	// 如果你只是想判断某个 key 是否存在而不关心它对应的值到底是多少，你可以这么做：

	// _, ok := map1[key1] // 如果key1存在则ok == true，否则ok为false

	// 	从 map1 中删除 key1：

	// 直接 delete(map1, key1) 就可以。

	// 如果 key1 不存在，该操作不会产生错误。

	var value int
	var isPresent bool

	map1 := make(map[string]int)
	map1["New Delhi"] = 55
	map1["Beijing"] = 20
	map1["Washington"] = 25
	value, isPresent = map1["Beijing"]
	if isPresent {
		fmt.Printf("The value of \"Beijing\" in map1 is: %d\n", value)
	} else {
		fmt.Printf("map1 does not contain Beijing")
	}

	value, isPresent = map1["Paris"]
	fmt.Printf("Is \"Paris\" in map1 ?: %t\n", isPresent)
	fmt.Printf("Value is: %d\n", value)

	// delete an item:
	delete(map1, "Washington")
	value, isPresent = map1["Washington"]
	if isPresent {
		fmt.Printf("The value of \"Washington\" in map1 is: %d\n", value)
	} else {
		fmt.Println("map1 does not contain Washington")
	}

}
