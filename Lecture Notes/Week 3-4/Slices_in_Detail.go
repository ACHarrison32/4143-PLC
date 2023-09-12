/*

package main

import "fmt"

func main() {
	var s []int

	t := []int{}
	u := make([]int, 5)
	v := make([]int, 0, 5)

	fmt.Printf("%d, %d, %T, %5t, %#[3]v\n", len(s), cap(s), s, s == nil)
	fmt.Printf("%d, %d, %T, %5t, %#[3]v\n", len(t), cap(t), t, t == nil)
	fmt.Printf("%d, %d, %T, %5t, %#[3]v\n", len(u), cap(u), u, u == nil)
	fmt.Printf("%d, %d, %T, %5t, %#[3]v\n", len(v), cap(v), v, v == nil)

}

*/

// ===============================================================================================================================

/*

package main

import "fmt"

func main() {
	a := [3]int{1, 2, 3}
	b := a[:1]

	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	c := b[0:2] //WTF

	fmt.Println("c = ", c)

	fmt.Println("len of b = ", len(b))
	fmt.Println("cap of b = ", cap(b))
	fmt.Println("len of c = ", len(c))
	fmt.Println("cap of c = ", cap(c))

	d := a[0:1:1] // [i:j:k] len j-i cap k-i

	fmt.Println("len of d = ", len(d))
	fmt.Println("cap of d = ", cap(d))
}

*/

package main

import "fmt"

func main() {
	a := [3]int{1, 2, 3}
	b := a[:1]
	c := b[0:2:2] //WTF

	fmt.Printf("a[%p] = %v\n", &a, a)
	fmt.Printf("b[%p] = %[1]v\n", b)
	fmt.Printf("c[%p] = %[1]v\n", c)

	c = append(c, 5)
	fmt.Printf("a[%p] = %v\n", &a, a)
	fmt.Printf("c[%p] = %[1]v\n", c)

	c[0] = 9
	fmt.Printf("a[%p] = %v\n", &a, a)
	fmt.Printf("c[%p] = %[1]v\n", c)
}
