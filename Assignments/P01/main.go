// This is the package declaration for the main program. Every Go program starts running in the main package.
package main

// This section imports external libraries and other packages that this file depends on.
import (
    // This imports the standard library's "fmt" package, which provides functions for formatted I/O.
	"fmt"

    // This imports a custom package from the path "example.com/program-1/mascot". This means that
    // the main.go program can access exported functions, types, etc. from the "mascot" package.
	"example.com/program-1/mascot"
)

// This is the entry point of the Go program. When this program is run, execution will start here.
func main() {
    // This line uses the "Println" function from the "fmt" package to print a message to the console.
    // It's calling the "BestMascot()" function from the imported "mascot" package to get a string
    // representation of the best mascot and then print it.
	fmt.Println(mascot.BestMascot())
}
