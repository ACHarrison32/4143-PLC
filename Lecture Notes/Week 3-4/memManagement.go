/*

Memory allocation and Deallocation happens automatically

new()
Allocate memort but no INIT
you will get a memory address
zeroed storage

make()
Allocate memory and INIT
you will get a memory address
non-zeroed storage

Garbage collection happens automatically
Out of scope or nil

https://pkg.go.dev/runtime => package runtime

The GOGC variable sets the initial garbage collection target percentage. A collection is
triggered when the ratio of freshly allocated data to live data remaining after the previous
collection reaches this percentage.

*/

