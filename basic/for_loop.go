package main

import "fmt"

func main() {
	// simple  iteration over a range
	count := 10
	for i := 0; i < count; i++ {
		fmt.Println(i)
	}

	// iterate over collection
	c := []int{1, 2, 3, 4, 5}
	for index, value := range c {
		fmt.Printf("Index: %v, Value: %v\n", index, value)
	}

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue //顾名思义 continue the loop but skip the rest of lines/statements
		}
		fmt.Println("Odd number: ", i)
		if i == 5 {
			break // break out of the loop
		}
	}

	rows := 5
	for i := 1; i < rows-i; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(" ")
		}
		for k := 1; k < 2*i-1; k++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}
