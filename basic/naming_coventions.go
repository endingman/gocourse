package main

import "fmt"

type UserInfo struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	user := UserInfo{"John", "Smith", 30}
	fmt.Println(user)

	// go 作用域内变量统一使用mixedCase 即首单词字母小写后续单词首字母大写
	var userInfoId = 100001
	fmt.Println(userInfoId)
}

// 命名规范 (场景)
//  PascalCase eg. UserInfo 所有单词首字母大写
//  MixedCase eg. userInfo 每个单词首字母小写
//  snake_case eg. user_info 下划线分隔
//  kebab-case eg. user-info 中划线分隔
//  uppercase eg. USER_INFO 常量名大写

//  go 文件命名统一使用snake_case
//  常量命名统一大写 eg. USER_INFO
