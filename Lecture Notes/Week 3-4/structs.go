/*

package main

import (
	"fmt"
	"time"
)

type Employee struct {
	Name   string
	Number int
	Boss   *Employee
	Hired  time.Time
}

func main() {
	c := map[string]*Employee{}

	c["Terry"] = &Employee{
		Name:   "Terry",
		Number: 2,
		Boss:   nil,
		Hired:  time.Now(),
	}

	c["Andrew"] = &Employee{
		Name:   "Andrew",
		Number: 1,
		Boss:   c["Terry"],
		Hired:  time.Now(),
	}

	//fmt.Printf("%T %[1]v\n", e)
	fmt.Printf("%T %+[1]v\n", c["Terry"])
	fmt.Printf("%T %+[1]v\n", c["Andrew"])

}

*/

// ====================================================================================================================================

/*

package main

import (
	"fmt"
)

type album1 struct {
	title  string
	artist string
	year   int
	copies int
}
type album2 = struct {
	title  string
	artist string
	year   int
	copies int
}

func main() {
	var a1 = album1{
		"The White Album",
		"The Beatles",
		1968,
		100000000000,
	}
	var a2 = album2{
		"The Banger Album",
		"Beanz",
		1999,
		200000000000,
	}

	fmt.Println(a1)
	fmt.Println(a2)

}

*/

// ====================================================================================================================================

/*

Two struct types are the compatible if
	the fields have a same type and name
	in the same roder
	and the same tages(*)

Structs are passed by value unless a pointer is used

dot notation works on pointers too, equivalent to (*a).copies

A struct with no fields is useful; it takes up no space

// a set type (instead of bool)
var isPresent map[int]struct{}

*/

// ====================================================================================================================================

/*

package main

import (
	"encoding/json"
	"fmt"
)

type Response struct {
	Page  int      `json:"page"`
	Words []string `json:"words, omitempty"`
}

func main() {
	r := Response{Page: 1, Words: []string{"up", "in", "out"}}

	j, _ := json.Marshal(r)
	fmt.Println(string(j))
	fmt.Printf("%#v\n", r)

	var r2 Response
	_ = json.Unmarshal(j, &r2)
	fmt.Printf("%#v\n", r2)

}

*/

// ====================================================================================================================================

/*

Tags can also be used in conjunction with SQL queries

import "github.vom/jmoiron/sqlx"

type item struct {
	Name string `db:"name"`
	When string `db:"created"`
}

func PutStats(db *sqlx.DB, item *item) error{
	stmt := `INSERT INTO items (name, created)
			 VALUES (:name, :created);`
	_, err := db.NamedExec(stmt, item)

	return err
}

*/