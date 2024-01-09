package main

import (
	"fmt"
)

const (
	logDeleted  = "user deleted"
	logNotFound = "user not found"
	logAdmin    = "admin deleted"
)

func logAndDelete(users map[string]user, name string) (log string) {
	defer delete(users, name)
	user, ok := users[name]
	if !ok {
		// maps are deleted by reference - if you delete a map, it'll be delete on the callers code as well
		// delete(users, name)
		return logNotFound
	}
	if user.admin {
		return logAdmin
	}
	// delete(users, name)
	return logDeleted
}

type user struct {
	name   string
	number int
	admin  bool
}

func test(users map[string]user, name string) {
	fmt.Printf("Attempting to delete %s...\n", name)
	defer fmt.Println("===========================")
	log := logAndDelete(users, name)
	fmt.Println("Log: ", log)
}
func main() {

	users := map[string]user{
		"john": {
			name:   "John",
			number: 42516614156,
			admin:  true,
		},
		"jane": {
			name:   "Jane",
			number: 42516614156,
			admin:  false,
		},
		"elon": {
			name:   "Elon",
			number: 12637738838,
			admin:  true,
		},
	}

	test(users)
}
