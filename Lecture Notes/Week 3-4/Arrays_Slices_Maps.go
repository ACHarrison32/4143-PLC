/*
All these are equivalent:

var a [3]int
var b [3]int{0, 0, 0}
var c [...]{0, 0, 0}  			//Sized by initializer

var d [3]int
d = b					//Elements copied

var m [...]int{1, 2, 3, 4}
c = m					//TYPE MISMATCH

Arrays are typed by sized, which is fixed at compile time
Arrays are passed by value, thus elements are copied.

*/

// ==================================================================================================================================

/*

var a []int						//nil, no storage
var b = []int{1, 2}					//initialized

a = append(a, 1)        				//append to nil OK
b = append(b, 3)					//[]int{1, 2, 3}

a = b 							//overwrites a

d := make([]int, 5) 					//[]int{0, 0, 0, 0, 0}
e := a							//same storage (alias)

e[0] == b[0]						//true

Slices have variable length, backed by some array;
Slices are passed by reference; no copying, updating OK

*/

// ==================================================================================================================================

/*

Slices are indexed like [8:11]
read as the starting element and one past the ending element, so this way we have 11-8=3 elemenets in our side

For loops work the same way in most cases:
	for i := 8; i < 11; i++ { 							//in math written [8, 11)
	...
	}

*/

// ==================================================================================================================================

/*

package main

import "fmt"

func main() {

	t := []byte("string")

	fmt.Println(len(t), t)          	 //What is t in this length
	fmt.Println(t[2])               	 //What is the third byte
	fmt.Println(t[:2])               	 //Two elements after third
	fmt.Println(t[2:])              	 //Everything after two
	fmt.Println(t[3:5])              	 //Fourth and fifth item
	fmt.Println(t[3:5], len(t[3:5])) 	 //What is the length of this

}

*/

// ==================================================================================================================================

// Slices Vs. Arrays

/*
Slices:
Variable length
Passed By Reference
Not Comparable
Cannot be used as map key
Has copy and append helpers
(Useful as function parameters)
*/
/*
Arrays:
Length fixed at compile time
Passed By Value (copied)
Comparable (==)
Can be used as map key
(Useful as "psuedo constants")

*/

//==================================================================================================================================

/*

var w = [...]int{1, 2, 3}				//Array of len(3)
var x = []int{0, 0, 0}					//Slice of len(3)

func do(a [3]int, b []int) []int {
	a = b 						//SYNTAX ERROR
	a[0] = 4 					//w unchanged
	b[0] = 3					//x changed

	c := make([]int, 5)				// copies only 3 elts
	c[4] = 42
	copy(c, b)					// copies only 3 elts

	return c
}

y:= do (w,x)
fmt.Printlconst(w, x, y)  				// [1 2 3] [3 0 0] [3 0 0 0 42]

*/

// ==================================================================================================================================

/*

Maps are dictionaries: ondexed by key, returning a value

You can read from a nil map, but insterting will panic

var m map[string]int					//nil, no storage
p := make(map[string]int)				//non-nil but empty

a := p["the"]						//returns 0
b := m["the"]						//same thing
m["and"] = 1						//PANIC - nil map
m = p
m["and"]++						//OK, same map as p now
c := p["and"]						//returns 1

mpas are passed by reference; no copying, updating OK
The type used for the key must ahve == and != defined (not slices, maps, or funcs)

Maps can't be comapred to one another; maps can be compared only to nil as a special case

var m =[string]int{
	"and": 1,
	"the": 1,
	"or": 2,
}

var n map[string]int

b := m == n						//SYNTAX ERROR
c := n == nil						//true
d := len(m)						//3
3 := cap(m)						//TYPE MISMATCH

*/

// ==================================================================================================================================

/*

len(s)							String						String length

len(a), cap(a)						Array						Array length, capacity (constant)

make(T, x)						Slice						Slice of type T with length x and capacity x
make(T, x, y)						Slice						Slice of type T with elngth x and capacity y

copy(c, d)						Slice						Copy from d to c; # = min of the two lengths
c = append(c, d)					Slice						append d to c and return a new slice result

len(s), cap(s)						Slice						Slice length and capacity

make(T)							Map						Map of type T
make(T, x)						Map						Map of type T with space hint for x elements

delete(m, k)						Map						Delete key k (if present, else no change)

len(m)							Map						Map length

*/

// ==================================================================================================================================

/*

Nil is a type of zero, it indicates the absence of something
Many Built ins are safe: len, cap, range

var s []int
var m map[string]int

l := len(s)						//length of nil slice is zero

i, ok := m["int"]					//0, false for any missing key

for _, v := range s{					//skip if s is nil or empty
	...
}

"make the zero value useful"

*/

// ==================================================================================================================================

/*

package main

import (
	"bufio"
	"fmt"
	"os"
	//"sort"
)

func main() {
	scan := bufio.NewScanner(os.Stdin)
	words := make(map[string]int)

	scan.Split(bufio.ScanWords)

	for scan.Scan() {
		words[scan.Text()]++
	}
	fmt.Println(len(words), "unique words")

}

*/

// ==================================================================================================================================

/*

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	scan := bufio.NewScanner(os.Stdin)
	words := make(map[string]int)

	scan.Split(bufio.ScanWords)

	for scan.Scan() {
		words[scan.Text()]++
	}
	fmt.Println(len(words), "unique words")

	type kv struct {
		key string
		val int
	}

	var ss []kv

	for k, v := range words {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].val > ss[j].val
	})

	for _, s := range ss[:3] {
		fmt.Println(s.key, "appears", s.val, "times")
	}

}

*/

// ==================================================================================================================================
