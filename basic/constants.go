package main

// const 常量 不可以改变
const (
	// Pi is a float
	Pi = 3.14

	GRAVITY = 9.81
)

func main() {
	const days int = 7

	//  const 块 把相关的的常量定义在一起，这样可以避免命名冲突 清晰易懂
	const (
		monday    = 1
		tuesday   = 2
		wednesday = 3
		thursday  = 4
		friday    = 5
		saturday  = 6
		sunday    = 7
	)
}
