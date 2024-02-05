/*
	Slices are only typed by elements they contain. An initialixed slice equals to nil and of 0 length
	To create empty slice with non-zero length use builtin `make` eg s=make([]int,10)
	By default an new slice capacity is equal to its length- check capacity using the builtin cap
	len() function - returns the length of the slice (the number of elements in the slice)
	cap() function - returns the capacity of the slice (the number of elements the slice can grow or shrink to)
	Slices support the following methods:
	- append - to add an element to the end of the slice
	- copy - creates a new slice for the given slice for copying
	- slice -

	Slices can be composed into multi-dimensional data structures. The length of the inner slices can vary
*/

package main

import (
	"fmt"
	"slices"
)

func main() {

	// uninitialized slice
	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	// non-zero length slice
	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	// set slice capacity explicitly
	s = make([]string, 3, 10)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	// append
	s = append(s, "a", "b", "c")
	fmt.Println("append:", s, "len:", len(s), "cap:", cap(s))

	// copy
	s = append(s, "d", "e", "f")
	fmt.Println("copy:", s, "len:", len(s), "cap:", cap(s))

	// slice
	s = append(s, "g", "h", "i")
	fmt.Println("slice:", s, "len:", len(s), "cap:", cap(s))

	/* setters and getters */
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	// slice lenght
	fmt.Println("len:", len(s))

	// append to a slice
	s = append(s, "n")
	s = append(s, "m", "k")
	fmt.Println("new slice:", s)

	// Slice copying
	c := make([]string, len(s))
	fmt.Println("copy:", c, "len:", len(c))
	copy(c, s)
	fmt.Println("copy:", c, "capacity:", cap(c))
	// append
	c = append(c, "new")
	c = append(c, "new2", "new3")
	fmt.Println("New copied and appended slice:", c, "capacity", cap(c))

	/* slice operator
	syntax
	slice[low:high]
	*/
	l := s[2:5] //this gets a slice of the elements s[2], s[3], and s[4].
	fmt.Println("sl2:", l)
	/* This slices up to (but excluding) s[5].
	And this slices up from (and including) s[2]. */

	// slice declaration/initialization shorthand
	t := []string{"j", "k", "l", "p"}
	fmt.Println("sl3:", t, "of len:", len(t), "and cap:", cap(t))

	// slcie funfunctions
	t2 := []string{"j", "k", "l", "p"}
	if slices.Equal(t, t2) {
		fmt.Println("t = t2")
	}

	// multi-dimensional slice
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		println("i:", i, "innerLen:", innerLen)
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
