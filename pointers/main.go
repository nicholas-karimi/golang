package main

import (
	"fmt"
	)

// When to use pointers
// 1. When you need to update a state.
// pointer = 8 bytes
// 2. When you want to optimize memory for large objects that are getting called ALOT.

// Why nit use *
// 1. Reduce memory allocation in the heap - Allocation is not free operation & less
// 2. Esacape analysis
type User struct {
	email string
	username string
	age int
}

func (u User) Email() string {
	return u.email
}

func (u *User) updateEmail(email string) {
	u.email = email
}

func Email(user *User) string {
	return user.email
}

func main(){
    user := User {
    email: "nkarimi@linux.com",	
    }
    user.updateEmail("nkarimi@google.com")
    fmt.Println(user.Email())
}

