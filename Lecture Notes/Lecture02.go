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

// package main

// import "fmt"

// func main() {
// 	// Basic for loop
// 	for i := 1; i <= 20; i++ {
// 		fmt.Print(i, " ")
// 	}
// 	fmt.Println()

// 	// Looping through an array/slice
// 	//numbers := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
// 	numbers2 := [5]int{0}

// 	numbers2[0] = 99
// 	numbers2[1] = 80
// 	numbers2[2] = 70
// 	for _, num := range numbers2 {
// 		fmt.Print(num, " ")
// 	}
// 	fmt.Println()
// }

// ======================================================

// package main

// import (
// 	"fmt"
// 	"slices"
// )

// func main() {

// 	var s []string
// 	fmt.Println("uninit:", s, s == nil, len(s) == 0)

// 	s = make([]string, 3)
// 	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

// 	s[0] = "a"
// 	s[1] = "b"
// 	s[2] = "c"
// 	fmt.Println("set:", s)
// 	fmt.Println("get:", s[2])

// 	fmt.Println("len:", len(s))

// 	s = append(s, "d")
// 	s = append(s, "e", "f")
// 	fmt.Println("apd:", s)

// 	c := make([]string, len(s))
// 	copy(c, s)
// 	fmt.Println("cpy:", c)

// 	l := s[2:5]
// 	fmt.Println("sl1:", l)

// 	l = s[:5]
// 	fmt.Println("sl2:", l)

// 	l = s[2:]
// 	fmt.Println("sl3:", l)

// 	t := []string{"g", "h", "i"}
// 	fmt.Println("dcl:", t)

// 	t2 := []string{"g", "h", "i"}
// 	if slices.Equal(t, t2) {
// 		fmt.Println("t == t2")
// 	}

// 	twoD := make([][]int, 3)
// 	for i := 0; i < 3; i++ {
// 		innerLen := i + 1
// 		twoD[i] = make([]int, innerLen)
// 		for j := 0; j < innerLen; j++ {
// 			twoD[i][j] = i + j
// 		}
// 	}
// 	fmt.Println("2d: ", twoD)
// }

// ======================================================

package main

import "fmt"

type Person struct {
	Name string
}

func (p Person) greet() string {
	return fmt.Sprintf("Hello, my name is %s.", p.Name)
}

type Calculator struct{}

func (c Calculator) square(number int) int {
	return number * number
}

func main() {
	person := Person{Name: "Andrew"}
	fmt.Println(person.greet())

	calculator := Calculator{}
	num := 12
	square := calculator.square(num)
	fmt.Printf("The square of %d is %d\n", num, square)
}
