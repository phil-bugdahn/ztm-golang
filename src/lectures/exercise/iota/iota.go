//--Summary:
//  Create a calculator that can perform basic mathematical operations.
//
//--Requirements:
//* Mathematical operations must be defined as constants using iota
//* Write a receiver function that performs the mathematical operation
//  on two operands
//* Operations required:
//  - Add, Subtract, Multiply, Divide
//* The existing function calls in main() represent the API and cannot be changed
//
//--Notes:
//* Your program is complete when it compiles and prints the correct results

package main

import "fmt"

type Operation byte
const (
	Add Operation = iota
	Subtract
	Multiply
	Divide
)

func (operation Operation) calculate(x int, y int) int {
	switch operation {
	case Add:
		return x + y
	case Subtract:
		return x - y
	case Multiply:
		return x * y
	case Divide:
		return x / y
	default:
		return 0
	}
}

func main() {
	add := Add
	sub := Operation(Subtract) // If Operation not specified in const then Operation(val) is necessary
	mul := Operation(Multiply)
	div := Operation(Divide)

	fmt.Println(add.calculate(2, 2)) // = 4

	fmt.Println(sub.calculate(10, 3)) // = 7

	fmt.Println(mul.calculate(3, 3)) // = 9

	fmt.Println(div.calculate(100, 2)) // = 50
}
