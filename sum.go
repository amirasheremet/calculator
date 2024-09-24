// the sum of two numbers
package main

import "fmt"

func main() {
	var num1, num2 int
	fmt.Print("Enter first number: ")
	fmt.Scanln(&num1)
	fmt.Print("Enter second number: ")
	fmt.Scanln(&num2)
	sum := num1 + num2
	fmt.Printf("The sum of %d and %d is %d\n", num1, num2, sum)
}
