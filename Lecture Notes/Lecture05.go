// package main

// import "fmt"

// func main() {
// 	var num int
// 	fmt.Printf("Enter a number: ")
// 	fmt.Scan(&num)
// 	fmt.Printf("You just entered: %v", num)
// }

/*

&num is used to pass the address of the variable num to the fmt.Scan() function. This is necessary because fmt.Scan()
expects a pointer as an argument, so it knows where to store the input value that the user provides.

The reason you use a pointer (&num) is so that when fmt.Scan() reads the input, it can directly modify the value of num
in memory. This way, after the function call, the value of num reflects the number that the user entered.

If you didn't use the & symbol to pass the address, the fmt.Scan() function wouldn't know where to store the user's input,
making the num variable unchanged and defeating the purpose of the scan.

So, to sum it up, &num is used to pass the memory address of the variable num so that fmt.Scan() can directly modify its
value based on user input.

*/

// ===========================================================================================================================================

/*

Basic Types:
%v: The value in a default format. Works for multiple types.
%T: The type of the value.
%%: A literal percent sign.

*/

// ===========================================================================================================================================

/*

Integer Types:
%b: Base 2 (binary).
%c: The character represented by the corresponding Unicode code point.
%d: Base 10 (decimal).
%o: Base 8 (octal).
%q: A single-quoted character literal safely escaped with Go syntax.
%x: Base 16 (hexadecimal), with lower-case letters.
%X: Base 16 (hexadecimal), with upper-case letters.
%U: Unicode format.

*/

// ===========================================================================================================================================

/*

Floating-point and Complex Types:
%e: Scientific notation, e.g., 1.234456e+78.
%E: Scientific notation, e.g., 1.234456E+78.
%f: Decimal with no exponent.
%F: Same as %f.
%g: %e or %f, whichever is shorter.
%G: %E or %F, whichever is shorter.

*/

// ===========================================================================================================================================

/*

String Types:
%s: Basic string output.
%q: A double-quoted string safely escaped with Go syntax.
%x: Base 16, lower-case, two characters per byte.
%X: Base 16, upper-case, two characters per byte.

*/

// ===========================================================================================================================================

/*

Pointer:
%p: Pointer address.

*/

// ===========================================================================================================================================

/*

package main

import "fmt"

func main() {
	count := 0
	for count < 5 {
		fmt.Println(count)
		count++
	}

	fruits := []string{"apple", "banana", "cherry"}
	for _, fruit := range fruits {
		fmt.Println(fruit)
	}
}

*/

// ===========================================================================================================================================

/*

In Go, the range keyword is used for iterating over elements in various types of collections, like slices, arrays,
maps, strings, and channels. When used in a for loop, range returns one or two values depending on the type of
collection you're iterating over. The first value is usually the index (or key for maps), and the second is the
value at that index (or the value for that key in maps).

*/