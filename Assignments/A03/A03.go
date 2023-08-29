// =================================================================
// Andrew Harrison
// Assignment 03 - "Hello World"
// 08-29-2023
// Professor Griffin
// 4143 PLC
// =================================================================

// =================================================================
// Assigment Description
// To figure out how to use GoLang and create a program that prints
// the string "Hello World"
// =================================================================

// =================================================================
// The main package is a special package in Go. It's not meant 
// to be imported by other Go programs; instead, it's the package 
// that contains the main function and is responsible for starting 
// and controlling the program's execution.
// =================================================================
package main

// =================================================================
// The import fmt statement is used to import the "fmt" package, 
// which stands for "format." The "fmt" package is part of the Go  
// standar library and provides functions for formatting input and 
// output. 
// =================================================================
import "fmt"

// =================================================================
// The main function. It serves as the entry point for your Go 
// program.
// =================================================================
func main() {

    // function from the "fmt" package, which stands for "format." 
    // It is used for printing text to the standard output.
    fmt.Println("Hello, World!")
}