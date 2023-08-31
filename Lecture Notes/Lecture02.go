// package main

// import "fmt"

// func main() {
// 	// Integer types
// 	var num int = 42
// 	var age = 30

// 	// Floating-point types
// 	var pi float32 = 3.14
// 	var height = 5.8

// 	// String type
// 	var message string = "Hello, Go!"

// 	// Boolean type
// 	var isTrue bool = false

// 	fmt.Println("num:", num)
// 	fmt.Println("age:", age)
// 	fmt.Println("pi:", pi)
// 	fmt.Println("height:", height)
// 	fmt.Println("message:", message)
// 	fmt.Println("isTrue:", isTrue)
// }

// ======================================================

// package main

// import "fmt"

// func main() {
// 	age := 18

// 	if age >= 18 {
// 		fmt.Println("You are an adult.")
// 	} else {
// 		fmt.Println("You are a minor.")
// 	}
// }

// ======================================================

// package main

// import "fmt"

// func main() {
// 	day := "Friday"

// 	switch day {
// 	case "Monday":
// 		fmt.Println("It's Monday.")
// 	case "Tuesday":
// 		fmt.Println("It's Tuesday.")
// 	case "Wednesday":
// 		fmt.Println("It's Wednesday.")
// 	default:
// 		fmt.Println("It's another day.")
// 	}
// }

// ======================================================

package main

import "fmt"

func main() {
	// Basic for loop
	for i := 1; i <= 20; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	// Looping through an array/slice
	//numbers := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	numbers2 := [5]int{0}

	numbers2[0] = 99
	numbers2[1] = 80
	numbers2[2] = 70
	for _, num := range numbers2 {
		fmt.Print(num, " ")
	}
	fmt.Println()
}

// ======================================================
