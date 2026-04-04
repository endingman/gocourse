package main

import "fmt"

func main() {

	// switch statement, in Go (switch case default), in other languages (switch case break default)
	// only the case block will be executed
	// if you want to execute more than one block, you can use fallthrough

	// switch  expression{
	// case value1:
	// Code to be executed if expression is equal to value1
	// case value2:
	// Code to be executed if expression is equal to value2
	// default:
	// }

	fruit := "apple"
	switch fruit {
	case "apple":
		fmt.Println("You have selected apple")
	case "banana":
		fmt.Println("You have selected banana")
	default:
		fmt.Println("You have selected something else")
	}

	// 多个条件 Multiple conditions
	day := "monday"

	switch day {
	case "monday", "tuesday", "wednesday", "thursday", "friday":
		fmt.Println("Working day")
	case "saturday", "sunday":
		fmt.Println("Weekend")
	default:
		fmt.Println("Invalid day")
	}

	//  fallthrough
	number := 2
	switch {
	case number > 1:
		fmt.Println("Greater than 1")
		fallthrough
	case number == 2:
		fmt.Println("Equal to 2")
	default:
		fmt.Println("Invalid number")
	}

	checkType(42)
	checkType(3.14)
	checkType("Hello")
	checkType(true)

}

func checkType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Integer: %d\n", v)
	case float64:
		fmt.Printf("Float: %f\n", v)
	case string:
		fmt.Printf("String: %s\n", v)
	default:
		fmt.Printf("Unknown type: %T\n", v)
	}
}
