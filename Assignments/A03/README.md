# :earth_americas: Assignment 03 - "Hello World"
## :bookmark_tabs: Assignemt Description
### Figure out how to use GoLang to create a program to print out "Hello World"

## :floppy_disk: Step 1 - Installation 
#### First you start by downloading GoLang on their website. For Windows users like me I went [Here](https://go.dev/dl/).
#### After installing GoLang you want to head over to your file explorer and right click on This PC. Then click on "Properties"
#### Once you have done this, click on "Advance System Settings" and locate the enviorment variable in the "Advanced" Section.
#### In the User Variable section check to make sure the Path has been added.
## :keyboard: Step 2 - Using GoLang
#### Once everything is download you are going to want to create your go program.
#### Example: A03.go 
#### Once you have created your file you are going to want to add your code:

``` Go
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
```

#### Once you have created your code you will then go to your command terminal and run the command: go build A03.go
#### this will create another file in your folder that is named similar to your go file. In this case it was named A03.exe, This is how you know its working
#### Once you have done this you want to type in the next command: go run A03.go, Make sure you are running the .go file and not the .exe file
#### It will then print your output:
<img src = "https://github.com/ACHarrison32/4143-PLC/blob/main/Assignments/A03/HelloWorld.PNG" >

