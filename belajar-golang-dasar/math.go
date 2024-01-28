package main

import "fmt"

func main() {
	var (
		a = 10
		b = 10
		c = 10
		i = 10
	)
	// Augmenented assignment
	i += 10
	i += 15
	// Unary operator
	i ++ // i = i + 1
	fmt.Println(a * b + c)
	fmt.Println(i)
}