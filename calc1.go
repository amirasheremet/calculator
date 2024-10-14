package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage: calculator <number1> <operator> <number2>")
		return
	}

	num1, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Println("Invalid number 1:", err)
		return
	}

	operator := os.Args[2]

	num2, err := strconv.ParseFloat(os.Args[3], 64)
	if err != nil {
		fmt.Println("Invalid number 2:", err)
		return
	}

	var result float64

	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0 {
			fmt.Println("Cannot divide by zero")
			return
		}
		result = num1 / num2
	default:
		fmt.Println("Invalid operator:", operator)
		return
	}

	fmt.Println("Result:", result)
}
