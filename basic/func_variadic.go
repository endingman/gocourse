package main

func main() {
	// 如果函数的最后一个参数是采用 ...type 的形式，那么这个函数就可以处理一个变长的参数，这个长度可以为 0，这样的函数称为变长函数。

	// func myFunc(a, b, arg ...int) {}
	Greeting("hello:", "Joe", "Anna", "Eileen")

	statement, total := sum("The total is:", 1, 2, 3, 4, 5)
	println(statement, total)
}

func Greeting(prefix string, who ...string) {
	for _, name := range who {
		println(prefix, name)
	}
}

func sum(returnStr string, nums ...int) (string, int) {
	total := 0
	for _, num := range nums {
		total += num
	}
	return returnStr, total
}
