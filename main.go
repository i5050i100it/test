package main

import (
	"fmt"
)

func add(x, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}

func multiply(x, y int) int {
	return x * y
}

func divide(x, y int) int {
	return x / y
}

func main() {
	var x, y int
	var symb string

	fmt.Println("Enter x and y...")
	fmt.Scan(&x, &y)
	fmt.Println("Enter symbol (+ ; - ; * ; / ;)")
	fmt.Scan(&symb)

	switch symb {
	case "+":
		fmt.Println("x + y =", add(x, y))
	case "-":
		fmt.Println("x - y =", subtract(x, y))
	case "*":
		fmt.Println("x * y =", multiply(x, y))
	case "/":
		fmt.Println("x / y =", divide(x, y))
	default:
		fmt.Println("Enter (+ ; - ; * ; / ;)")
	}

}
