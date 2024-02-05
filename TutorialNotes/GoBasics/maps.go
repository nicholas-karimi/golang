/*
 Mapps are go builtin associative array data structure- also known as hash or dicts in other languages
	sytnax
	use builtin make
	make(map[key-type]value-type)
	 maps appear in the form map[k:v k:v] when printed with fmt.Println.
*/

package main

import (
	"fmt"
	"maps"
)

func main() {
	users := make(map[string]string)
	users["first_name"] = "John"
	users["last_name"] = "Doe"
	fmt.Println("users:", users)

	// get key value
	first_name := users["first_name"]
	fmt.Println("first name:", first_name)

	/* if the key is not found the zero value will be returned */
	email := users["email"]
	fmt.Println("email:", email)

	// get length of map
	fmt.Println("len:", len(users))

	// delete key
	delete(users, "last_name")
	fmt.Println("deleted users:", users)

	// // check if key exists
	_, found := users["last_name"]
	fmt.Println("last name:", found)

	// // check if map is empty
	if len(users) == 0 {
		fmt.Println("empty")
	} else {
		fmt.Println("not empty", users)
	}

	// shorthand for maps
	gadgets := map[string]int{
		"keyboard": 100,
		"mouse":    50,
	}
	fmt.Println("gadgets:", gadgets)
	// update
	gadgets["keyboard"] = 2000
	fmt.Println("updated gadgets:", gadgets)

	// add new item
	gadgets["monitor"] = 1000
	fmt.Println("updated gadgets:", gadgets)

	// methods
	gadgets2 := map[string]int{
		"keyboard": 2000,
		"mouse":    50,
		"monitor":  1000,
	}
	if maps.Equal(gadgets, gadgets2) {
		fmt.Println("equal dict")
	} else {
		fmt.Println("not equal")
	}

	// loop through
	for key, value := range gadgets {
		// fmt.Printf("key:", key, "value:", value)
		fmt.Printf("%s: %d, ", key, value)
	}

	// clear map
	clear(users)
	fmt.Println("cleared users:", users)

	// iterate maps in order
	new_dict := map[string]int{
		"keyboard": 2000,
		"mouse":    50,
		"monitor":  1000,
	}

	var nd []string // define the order
	nd = append(nd, "keyboard", "mouse", "monitor")

	for _, key := range nd {
		fmt.Printf("%s: %d, ", key, new_dict[key])
	}
}
