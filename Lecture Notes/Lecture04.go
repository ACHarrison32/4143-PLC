package main // Declare the main package, which is used for executable programs.

import (
	"fmt"  // Import the 'fmt' package for formatted input and output.
	"math" // Import the 'math' package for mathematical functions.
)

func main() { // The main function, which is the entry point of the program.

	var num1, num2 int64 // Declare two int64 variables 'num1' and 'num2'.
	var absolute int64   // Declare an int64 variable 'absolute' to store the absolute difference.

	for { // Start an infinite loop.
		fmt.Println("Enter Number 1: ")
		fmt.Scanf("%d", &num1) // Read two integers from standard input.
		fmt.Println("Enter Number 2: ")
		fmt.Scan("%d", &num2)

		// Calculate the absolute difference between 'num1' and 'num2'.
		absolute = int64(math.Abs(float64(num1 - num2)))

		fmt.Printf("The absolute value of %v and %v is %v", num1, num2, absolute)
	}

	// The program ends when the loop is exited due to an error or EOF.
}

/* This is a simple example of how to use a global variable in Python
A = 10

def func(){
	global A
}
*/

/*
Arrays: Fixed size sequence of elements of the same type  [size]type
Slices: Dynamic arrays with flexible size    []type
Maps: Key-value pairs (association arrays)   map[keyType]valueType
Structs: Composite data type that groups fields with different data types
*/

/*

 */
