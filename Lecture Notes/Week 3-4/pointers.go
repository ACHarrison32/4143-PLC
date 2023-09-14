/*

This program currently doesnt work

package main

import "fmt"

type Player struct {
	health int
}

func takeDamageExplosion(player Player) {
	fmt.Println("Player is taking damage from explosion")
	explosionDmg := 10
	player.health -= explosionDmg
}

func main() {
	player := Player{
		health: 100,
	}
	fmt.Printf("before explosion %+v\n", player)
	takeDamageExplosion(player)	// copy
	fmt.Printf("after explosion %+v\n", player)

}

*/

// ====================================================================================================================================

/*

package main

import "fmt"

type Player struct {
	health int
}

func takeDamageExplosion(player *Player) {
	fmt.Println("Player is taking damage from explosion")
	explosionDmg := 10
	player.health -= explosionDmg
}

func main() {
	player := &Player{
		health: 100,
	}
	// 8 byte long integer (pointer)

	fmt.Printf("Before explosion %+v\n", player)
	takeDamageExplosion(player) // copy, you could also do &player here instead of line 48
	fmt.Printf("After explosion %+v\n", player)

}

*/

// ====================================================================================================================================

/*

package main

import "fmt"

type Player struct {
	health int
}

// These next 2 functions are the exact same just written different

func (player *Player) takeDamageExplosion(dmg int) {
	fmt.Println("Player is taking damage from explosion")
	player.health -= dmg
}

func takeDamageExplosion(player *Player, dmg int) {
	fmt.Println("Player is taking damage from explosion")
	player.health -= dmg
}

func main() {
	player := &Player{
		health: 100,
	}
	// 8 byte long integer (pointer)

	fmt.Printf("Before explosion %+v\n", player)
	player.takeDamageExplosion(50)
	fmt.Printf("After explosion %+v\n", player)

}

*/

//====================================================================================================================================

/*

1. Pointers allow you to bypass the copying of a variable, what allows you to save memory,
as well as change the object by reference without copying it (this was already mentioned in the video)
2. The pointer value can be nil (var myPointer *int), so it is very convenient to use it,
for example, where we return an object from the database, but under certain conditions the object may
not be found in the database (the classic example is ORM):

type UserID int

type User struct {
    userId UserID
    firstName string
    lastName string
}

func FindUser(id UserID) *User {
    var userId UserId = 123

    if id == userId {
        return &User{
            userId: userId
            firstName: "First Name",
            lastName: "Last Name",
        }
    }
    return nil
}

func main() {
    if user := FindUser(UserID(123)); user != nil {
        fmt.Println(user.firstName + " " + user.lastName)
    }
}

3. Sometimes it is convenient to create a pointer to a particular type with a default value that matches
that type. To do this, you can use the "new" function, which returns a pointer to an initialized object
in memory with a default value:

{
    var myPointer *int // myPointer is nil
}
{
    myPointer := new(int) // myPointer is an address to an integer type with value 0
}

You need to be careful by using pointers, because the uncontrolled use of pointers puts a load on
the garbage collector.

*/