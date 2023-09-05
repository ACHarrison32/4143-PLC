package main

import "fmt"

func main() {

	// strings
	var nameOne string = "Mario"
	var nameTwo = "Luigi"
	var nameThree string

	fmt.Println(nameOne)
	fmt.Println(nameTwo)
	fmt.Println(nameThree)

	nameOne = "Peach"
	nameThree = "Bowser"

	fmt.Println(nameOne, nameTwo, nameThree)

	nameFour := "Yoshi"

	fmt.Println(nameOne, nameTwo, nameThree, nameFour)

	// Ints
	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40

	fmt.Println(ageOne, ageTwo, ageThree)

}
