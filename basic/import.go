package main

import (
	"fmt"
	_http "net/http" //引入别名
)

func main() {
	fmt.Println("hello, go standard library")
	resp, err := _http.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		fmt.Println("Err:", err)
	}

	defer resp.Body.Close()

	fmt.Println("Http Response Status:", resp.Status)
}
