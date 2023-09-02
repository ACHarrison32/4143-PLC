// ===================================================================================================== EXAMPLE ONE

// package main

// import "fmt"

// func main() {
// 	fmt.Println("Hello World")
// }

// ===================================================================================================== VARIABLES AND CONSTANTS

// package main

// import "fmt"

// func main() {

// 	var conferenceName = "Go Conference"
// 	const conferencetTickets = 50
// 	var remainingTickets = 50

// 	fmt.Println("Welcome to", conferenceName, "booking applicaiton")
// 	// fmt.Printf("Welcome to" %v "booking applicaiton\n", conferenceName)
// 	fmt.Println("We have a total of", conferencetTickets, "tickets and", remainingTickets,
// "are still available")
// 	// fmt.Printf("We have a total of" %v "tickets and" %v "are still available\n",
// conferenceTickets, remainingTickets)
// 	fmt.Println("Get your tickets here to attend")
// }

// ===================================================================================================== DATA TYPES
// You need to tell Go Compiler, the data type when declaring the variable
// Type Inference: BUT, Go can infer the type when you assign a value

// package main

// import "fmt"

// func main() {

// 	conferenceName := "Go Conference"
// 	const conferencetTickets = 50
// 	var remainingTickets = 50

// 	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n",
//  conferencetTickets, remainingTickets, conferenceName)

// 	fmt.Println("Welcome to", conferenceName, "booking applicaiton")
// 	// fmt.Printf("Welcome to" %v "booking applicaiton\n", conferenceName)
// 	fmt.Println("We have a total of", conferencetTickets, "tickets and", remainingTickets,
//  "are still available")
// 	// fmt.Printf("We have a total of" %v "tickets and" %v "are still available\n",
// conferenceTickets, remainingTickets)
// 	fmt.Println("Get your tickets here to attend")

// 	var userName string
// 	var userTickets int
// 	// ask user for their name
// 	fmt.Println("Enter your first name: ")
// 	fmt.Scan(&userName)

// 	userTickets = 2
// 	fmt.Println("User", userName, "booked", userTickets, "tickets.")
// }

// ===================================================================================================== GETTING USER INPUT DATA

// package main

// import "fmt"

// func main() {

// 	conferenceName := "Go Conference"
// 	const conferencetTickets = 50
// 	var remainingTickets = 50

// 	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n",
//  conferencetTickets, remainingTickets, conferenceName)

// 	fmt.Println("Welcome to", conferenceName, "booking applicaiton")
// 	// fmt.Printf("Welcome to" %v "booking applicaiton\n", conferenceName)
// 	fmt.Println("We have a total of", conferencetTickets, "tickets and", remainingTickets,
//  "are still available")
// 	// fmt.Printf("We have a total of" %v "tickets and" %v "are still available\n",
// conferenceTickets, remainingTickets)
// 	fmt.Println("Get your tickets here to attend")

// 	var firstName string
// 	var lastName string
// 	var email string
// 	var userTickets int
// 	// ask user for their name
// 	fmt.Println("Enter your first name: ")
// 	fmt.Scan(&firstName)
// 	fmt.Println("Enter your last name: ")
// 	fmt.Scan(&lastName)
// 	fmt.Println("Enter your email address: ")
// 	fmt.Scan(&email)
// 	fmt.Println("Enter number of tickets: ")
// 	fmt.Scan(&userTickets)

// 	fmt.Println("Thank you", firstName, lastName, "for booking", userTickets, "tickets.
//  You will reviece a confirmation email at", email)
// }

// ===================================================================================================== BOOK TICKET LOGIC

// package main

// import "fmt"

// func main() {

// 	conferenceName := "Go Conference"
// 	const conferencetTickets = 50
// 	var remainingTickets uint = 50

// 	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n",
// 		conferencetTickets, remainingTickets, conferenceName)

// 	fmt.Println("Welcome to", conferenceName, "booking applicaiton")
// 	// fmt.Printf("Welcome to" %v "booking applicaiton\n", conferenceName)
// 	fmt.Println("We have a total of", conferencetTickets, "tickets and", remainingTickets,
// 		"are still available")
// 	// fmt.Printf("We have a total of" %v "tickets and" %v "are still available\n",
// 	// conferenceTickets, remainingTickets)
// 	fmt.Println("Get your tickets here to attend")

// 	var firstName string
// 	var lastName string
// 	var email string
// 	var userTickets uint
// 	// ask user for their name
// 	fmt.Println("Enter your first name: ")
// 	fmt.Scan(&firstName)
// 	fmt.Println("Enter your last name: ")
// 	fmt.Scan(&lastName)
// 	fmt.Println("Enter your email address: ")
// 	fmt.Scan(&email)
// 	fmt.Println("Enter number of tickets: ")
// 	fmt.Scan(&userTickets)

// 	remainingTickets = remainingTickets - userTickets

// 	fmt.Println("Thank you", firstName, lastName, "for booking", userTickets,
// 		"tickets. You will reviece a confirmation email at", email)
// 	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
// }

// ===================================================================================================== ARRAYS AND SLICES IN GO
// Slice is an abstraction of an array
// Slices are more flexible and powerful: variable-length or get a sub-array of its own
// Slices are also index-based and have a size, but is resized when needed
// Append adds the elements at the end of the slice
// Append grows the slice if a greater capacity is needed and returns the updated slice value

// package main

// import "fmt"

// func main() {

// 	conferenceName := "Go Conference"
// 	const conferencetTickets = 50
// 	var remainingTickets uint = 50

// 	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n",
// 		conferencetTickets, remainingTickets, conferenceName)

// 	fmt.Println("Welcome to", conferenceName, "booking applicaiton")
// 	// fmt.Printf("Welcome to" %v "booking applicaiton\n", conferenceName)
// 	fmt.Println("We have a total of", conferencetTickets, "tickets and", remainingTickets,
// 		"are still available")
// 	// fmt.Printf("We have a total of" %v "tickets and" %v "are still available\n",
// 	// conferenceTickets, remainingTickets)
// 	fmt.Println("Get your tickets here to attend")

// 	var bookings []string
// 	// bookings := []string{} Same as the line above

// 	var firstName string
// 	var lastName string
// 	var email string
// 	var userTickets uint
// 	// ask user for their name
// 	fmt.Println("Enter your first name: ")
// 	fmt.Scan(&firstName)
// 	fmt.Println("Enter your last name: ")
// 	fmt.Scan(&lastName)
// 	fmt.Println("Enter your email address: ")
// 	fmt.Scan(&email)
// 	fmt.Println("Enter number of tickets: ")
// 	fmt.Scan(&userTickets)

// 	remainingTickets = remainingTickets - userTickets
// 	bookings = append(bookings, firstName+" "+lastName)

// 	//fmt.Printf("The whole array: %v\n", bookings)
// 	//fmt.Printf("The whole slice: %v\n", bookings)
// 	//fmt.Printf("The first value: %v\n", bookings[0])
// 	//fmt.Printf("The array type: %T\n", bookings)
// 	//fmt.Printf("The slice type: %T\n", bookings)
// 	//fmt.Printf("The array length: %v\n", len(bookings))
// 	//fmt.Printf("The slice length: %v\n", len(bookings))

// 	fmt.Println("Thank you", firstName, lastName, "for booking", userTickets,
// 		"tickets. You will reviece a confirmation email at", email)
// 	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

// 	fmt.Printf("These are all our bookings: %v\n", bookings)
// }

// ===================================================================================================== Loops
// You only have the "for loop" loop
// Range iterates over elements for different data strictures (so not only arrays and slices) (line 279)
// string.Fields() splits the string with white space as seperator (line 280)
// _ Blank identifier to ignore a variable you dont want to use (line 279)
// with Go you need to make unused variables explicit

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {

// 	conferenceName := "Go Conference"
// 	const conferencetTickets = 50
// 	var remainingTickets uint = 50
// 	var bookings []string
// 	// bookings := []string{} Same as the line above

// 	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n",
// 		conferencetTickets, remainingTickets, conferenceName)

// 	fmt.Println("Welcome to", conferenceName, "booking applicaiton")
// 	// fmt.Printf("Welcome to" %v "booking applicaiton\n", conferenceName)
// 	fmt.Println("We have a total of", conferencetTickets, "tickets and", remainingTickets,
// 		"are still available")
// 	// fmt.Printf("We have a total of" %v "tickets and" %v "are still available\n",
// 	// conferenceTickets, remainingTickets)
// 	fmt.Println("Get your tickets here to attend")

// 	for {
// 		var firstName string
// 		var lastName string
// 		var email string
// 		var userTickets uint

// 		fmt.Println("Enter your first name: ")
// 		fmt.Scan(&firstName)
// 		fmt.Println("Enter your last name: ")
// 		fmt.Scan(&lastName)
// 		fmt.Println("Enter your email address: ")
// 		fmt.Scan(&email)
// 		fmt.Println("Enter number of tickets: ")
// 		fmt.Scan(&userTickets)

// 		remainingTickets = remainingTickets - userTickets
// 		bookings = append(bookings, firstName+" "+lastName)

// 		//fmt.Printf("The whole array: %v\n", bookings)
// 		//fmt.Printf("The whole slice: %v\n", bookings)
// 		//fmt.Printf("The first value: %v\n", bookings[0])
// 		//fmt.Printf("The array type: %T\n", bookings)
// 		//fmt.Printf("The slice type: %T\n", bookings)
// 		//fmt.Printf("The array length: %v\n", len(bookings))
// 		//fmt.Printf("The slice length: %v\n", len(bookings))

// 		fmt.Println("Thank you", firstName, lastName, "for booking", userTickets,
// 			"tickets. You will reviece a confirmation email at", email)
// 		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

// 		firstNames := []string{}
// 		for _, booking := range bookings {
// 			var names = strings.Fields(booking)
// 			firstNames = append(firstNames, names[0])
// 		}
// 		fmt.Printf("These are all our bookings: %v\n", firstNames)
// 	}
// }

// ===================================================================================================== If Else Statements
// Break statement terminates the for loop and continues with the code right after the for loop
// Continue statement causes loop to skip the remainder of its body and immediately retesting its condition

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {

// 	conferenceName := "Go Conference"
// 	const conferencetTickets = 50
// 	var remainingTickets uint = 50
// 	var bookings []string
// 	// bookings := []string{} Same as the line above

// 	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n",
// 		conferencetTickets, remainingTickets, conferenceName)

// 	fmt.Println("Welcome to", conferenceName, "booking applicaiton")
// 	// fmt.Printf("Welcome to" %v "booking applicaiton\n", conferenceName)
// 	fmt.Println("We have a total of", conferencetTickets, "tickets and", remainingTickets,
// 		"are still available")
// 	// fmt.Printf("We have a total of" %v "tickets and" %v "are still available\n",
// 	// conferenceTickets, remainingTickets)
// 	fmt.Println("Get your tickets here to attend")

// 	for {
// ------------------------------------------------------------------------------------------------ If Statement
// var firstName string
// var lastName string
// var email string
// var userTickets uint

// fmt.Println("Enter your first name: ")
// fmt.Scan(&firstName)
// fmt.Println("Enter your last name: ")
// fmt.Scan(&lastName)
// fmt.Println("Enter your email address: ")
// fmt.Scan(&email)
// fmt.Println("Enter number of tickets: ")
// fmt.Scan(&userTickets)

// if userTickets > remainingTickets {
// 	fmt.Printf("We only have %v tickets remaining, so you cant book %v tickets dumbass\n",
// remainingTickets, userTickets)
// 	continue
// }

// remainingTickets = remainingTickets - userTickets
// bookings = append(bookings, firstName+" "+lastName)

// //fmt.Printf("The whole array: %v\n", bookings)
// //fmt.Printf("The whole slice: %v\n", bookings)
// //fmt.Printf("The first value: %v\n", bookings[0])
// //fmt.Printf("The array type: %T\n", bookings)
// //fmt.Printf("The slice type: %T\n", bookings)
// //fmt.Printf("The array length: %v\n", len(bookings))
// //fmt.Printf("The slice length: %v\n", len(bookings))

// fmt.Println("Thank you", firstName, lastName, "for booking", userTickets,
// 	"tickets. You will reviece a confirmation email at", email)
// fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

// firstNames := []string{}
// for _, booking := range bookings {
// 	var names = strings.Fields(booking)
// 	firstNames = append(firstNames, names[0])
// }
// fmt.Printf("These are all our bookings: %v\n", firstNames)

// if remainingTickets == 0 {
// 	// end program
// 	fmt.Println("Our conference is booked up. Sorry bout ya luck")
// 	break
// }

// ------------------------------------------------------------------------------------------------ If Else Statement
// Or you can do something like this for an is else statement:
// 		var firstName string
// 		var lastName string
// 		var email string
// 		var userTickets uint

// 		fmt.Println("Enter your first name: ")
// 		fmt.Scan(&firstName)
// 		fmt.Println("Enter your last name: ")
// 		fmt.Scan(&lastName)
// 		fmt.Println("Enter your email address: ")
// 		fmt.Scan(&email)
// 		fmt.Println("Enter number of tickets: ")
// 		fmt.Scan(&userTickets)

// 		if userTickets <= remainingTickets {
// 			remainingTickets = remainingTickets - userTickets
// 			bookings = append(bookings, firstName+" "+lastName)

// 			//fmt.Printf("The whole array: %v\n", bookings)
// 			//fmt.Printf("The whole slice: %v\n", bookings)
// 			//fmt.Printf("The first value: %v\n", bookings[0])
// 			//fmt.Printf("The array type: %T\n", bookings)
// 			//fmt.Printf("The slice type: %T\n", bookings)
// 			//fmt.Printf("The array length: %v\n", len(bookings))
// 			//fmt.Printf("The slice length: %v\n", len(bookings))

// 			fmt.Println("Thank you", firstName, lastName, "for booking", userTickets,
// 				"tickets. You will reviece a confirmation email at", email)
// 			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

// 			firstNames := []string{}
// 			for _, booking := range bookings {
// 				var names = strings.Fields(booking)
// 				firstNames = append(firstNames, names[0])
// 			}
// 			fmt.Printf("These are all our bookings: %v\n", firstNames)

// 			if remainingTickets == 0 {
// 				// end program
// 				fmt.Println("Our conference is booked up. Sorry bout ya luck")
// 				break
// 			}
// 		} else {
// 			fmt.Printf("We only have %v tickets remaining, so you cant book %v tickets dumbass\n",
// 		remainingTickets, userTickets)
// 		}
// 	}
// }

// You can also do this with the for loop
// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {

// 	conferenceName := "Go Conference"
// 	const conferencetTickets = 50
// 	var remainingTickets uint = 50
// 	var bookings []string
// 	// bookings := []string{} Same as the line above

// 	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n",
// 		conferencetTickets, remainingTickets, conferenceName)

// 	fmt.Println("Welcome to", conferenceName, "booking applicaiton")
// 	// fmt.Printf("Welcome to" %v "booking applicaiton\n", conferenceName)
// 	fmt.Println("We have a total of", conferencetTickets, "tickets and", remainingTickets,
// 		"are still available")
// 	// fmt.Printf("We have a total of" %v "tickets and" %v "are still available\n",
// 	// conferenceTickets, remainingTickets)
// 	fmt.Println("Get your tickets here to attend")

// 	for remainingTickets > 0 && len(bookings)< 50{
// 		var firstName string
// 		var lastName string
// 		var email string
// 		var userTickets uint

// 		fmt.Println("Enter your first name: ")
// 		fmt.Scan(&firstName)
// 		fmt.Println("Enter your last name: ")
// 		fmt.Scan(&lastName)
// 		fmt.Println("Enter your email address: ")
// 		fmt.Scan(&email)
// 		fmt.Println("Enter number of tickets: ")
// 		fmt.Scan(&userTickets)

// 		if userTickets <= remainingTickets {
// 			remainingTickets = remainingTickets - userTickets
// 			bookings = append(bookings, firstName+" "+lastName)

// 			//fmt.Printf("The whole array: %v\n", bookings)
// 			//fmt.Printf("The whole slice: %v\n", bookings)
// 			//fmt.Printf("The first value: %v\n", bookings[0])
// 			//fmt.Printf("The array type: %T\n", bookings)
// 			//fmt.Printf("The slice type: %T\n", bookings)
// 			//fmt.Printf("The array length: %v\n", len(bookings))
// 			//fmt.Printf("The slice length: %v\n", len(bookings))

// 			fmt.Println("Thank you", firstName, lastName, "for booking", userTickets,
// 				"tickets. You will reviece a confirmation email at", email)
// 			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

// 			firstNames := []string{}
// 			for _, booking := range bookings {
// 				var names = strings.Fields(booking)
// 				firstNames = append(firstNames, names[0])
// 			}
// 			fmt.Printf("These are all our bookings: %v\n", firstNames)

// 			if remainingTickets == 0 {
// 				// end program
// 				fmt.Println("Our conference is booked up. Sorry bout ya luck")
// 				break
// 			}
// 		} else {
// 			fmt.Printf("We only have %v tickets remaining, so you cant book %v tickets dumbass\n",
// 		remainingTickets, userTickets)
// 		}
// 	}
// }

// ===================================================================================================== Validating User Input

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {

// 	conferenceName := "Go Conference"
// 	const conferencetTickets = 50
// 	var remainingTickets uint = 50
// 	var bookings []string
// 	// bookings := []string{} Same as the line above

// 	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n",
// 		conferencetTickets, remainingTickets, conferenceName)

// 	fmt.Println("Welcome to", conferenceName, "booking applicaiton")
// 	// fmt.Printf("Welcome to" %v "booking applicaiton\n", conferenceName)
// 	fmt.Println("We have a total of", conferencetTickets, "tickets and", remainingTickets,
// 		"are still available")
// 	// fmt.Printf("We have a total of" %v "tickets and" %v "are still available\n",
// 	// conferenceTickets, remainingTickets)
// 	fmt.Println("Get your tickets here to attend")

// 	for remainingTickets > 0 && len(bookings) < 50 {
// 		var firstName string
// 		var lastName string
// 		var email string
// 		var userTickets uint

// 		fmt.Println("Enter your first name: ")
// 		fmt.Scan(&firstName)
// 		fmt.Println("Enter your last name: ")
// 		fmt.Scan(&lastName)
// 		fmt.Println("Enter your email address: ")
// 		fmt.Scan(&email)
// 		fmt.Println("Enter number of tickets: ")
// 		fmt.Scan(&userTickets)

// 		var isValidName bool = len(firstName) >= 2 && len(lastName) >= 2
// 		var isValidEmail bool = strings.Contains(email, "@")
// 		var isValidTicketNumber = userTickets > 0 && userTickets <= remainingTickets

// 		if isValidName && isValidEmail && isValidTicketNumber {
// 			remainingTickets = remainingTickets - userTickets
// 			bookings = append(bookings, firstName+" "+lastName)

// 			//fmt.Printf("The whole array: %v\n", bookings)
// 			//fmt.Printf("The whole slice: %v\n", bookings)
// 			//fmt.Printf("The first value: %v\n", bookings[0])
// 			//fmt.Printf("The array type: %T\n", bookings)
// 			//fmt.Printf("The slice type: %T\n", bookings)
// 			//fmt.Printf("The array length: %v\n", len(bookings))
// 			//fmt.Printf("The slice length: %v\n", len(bookings))

// 			fmt.Println("Thank you", firstName, lastName, "for booking", userTickets,
// 				"tickets. You will reviece a confirmation email at", email)
// 			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

// 			firstNames := []string{}
// 			for _, booking := range bookings {
// 				var names = strings.Fields(booking)
// 				firstNames = append(firstNames, names[0])
// 			}
// 			fmt.Printf("These are all our bookings: %v\n", firstNames)

// 			if remainingTickets == 0 {
// 				// end program
// 				fmt.Println("Our conference is booked up. Sorry bout ya luck")
// 				break
// 			}
// 		} else {
// 			if !isValidName {
// 				fmt.Println("I know your lying about your name. Try again buckaroo!")
// 			}
// 			if !isValidEmail {
// 				fmt.Println("Aren't emails suppose to have and @ symbol. Where was yours? Try harder next time")
// 			}
// 			if !isValidTicketNumber {
// 				fmt.Printf("Now if we only have %v tickets remaining, how do you expect to recieve %v..."+
//					"Didnt think about that did ya?\n", remainingTickets, userTickets)
// 			}
// 		}
// 	}
// }

// ===================================================================================================== Switch Statements

package main

import (
	"fmt"
	"strings"
)

func main() {

	conferenceName := "Go Conference"
	const conferencetTickets = 50
	var remainingTickets uint = 50
	var bookings []string
	// bookings := []string{} Same as the line above

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n",
		conferencetTickets, remainingTickets, conferenceName)

	fmt.Println("Welcome to", conferenceName, "booking applicaiton")
	// fmt.Printf("Welcome to" %v "booking applicaiton\n", conferenceName)
	fmt.Println("We have a total of", conferencetTickets, "tickets and", remainingTickets,
		"are still available")
	// fmt.Printf("We have a total of" %v "tickets and" %v "are still available\n",
	// conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for remainingTickets > 0 && len(bookings) < 50 {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)
		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)
		fmt.Println("Enter your email address: ")
		fmt.Scan(&email)
		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		var isValidName bool = len(firstName) >= 2 && len(lastName) >= 2
		var isValidEmail bool = strings.Contains(email, "@")
		var isValidTicketNumber = userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			//fmt.Printf("The whole array: %v\n", bookings)
			//fmt.Printf("The whole slice: %v\n", bookings)
			//fmt.Printf("The first value: %v\n", bookings[0])
			//fmt.Printf("The array type: %T\n", bookings)
			//fmt.Printf("The slice type: %T\n", bookings)
			//fmt.Printf("The array length: %v\n", len(bookings))
			//fmt.Printf("The slice length: %v\n", len(bookings))

			fmt.Println("Thank you", firstName, lastName, "for booking", userTickets,
				"tickets. You will reviece a confirmation email at", email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("These are all our bookings: %v\n", firstNames)

			if remainingTickets == 0 {
				// end program
				fmt.Println("Our conference is booked up. Sorry bout ya luck")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("I know your lying about your name. Try again buckaroo!")
			}
			if !isValidEmail {
				fmt.Println("Aren't emails suppose to have and @ symbol. Where was yours? Try harder next time")
			}
			if !isValidTicketNumber {
				fmt.Printf("Now if we only have %v tickets remaining, how do you expect to recieve %v..."+
					"Didnt think about that did ya?\n", remainingTickets, userTickets)
			}
		}
	}
}
