package main

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io"
)

func getUserMap(names []string, phoneNumber []int) (map[string]user, error) {
	userMap := make(map[string]user)
	if len(names) != len(phoneNumber) {
		return nil, errors.New("invalid sizes")
	}
	for i := 0; i < len(names); i++ {
		name := names[i]
		phoneNumber := phoneNumber[i]
		userMap[name] = user{
			name:        name,
			phoneNumber: phoneNumber,
		}
	}
	return userMap, nil
}

type user struct {
	name        string
	phoneNumber int
}

func test(names []string, phoneNumber []int) {
	fmt.Println("Creating a map....")
	defer fmt.Println("=============================================")
	users, err := getUserMap(names, phoneNumber)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, name := range names {
		fmt.Printf("key: %v, value: \n", name)
		fmt.Println("- name: ", users[name].name)
		fmt.Println("- phoneNumber: ", users[name].phoneNumber)
	}
}

// count how many times user id appears in the slice to track how amny msgs received

func getCount(userIDS []string) map[string]int {
	counts := make(map[string]int)
	for _, userID := range userIDS {
		count := counts[userID]
		count++
		counts[userID] = count
	}
	return counts
}

func testCount(userIDS []string, ids []string ){
	fmt.Printf("Generating counts for %v user IDs..\n", len(userIDS))
	counts := getCount(userIDS)
	fmt.Println("Counts from select IDS:")
	for _, k := range ids {
		v := counts[k]
		fmt.Printf(" - %s: %d\n", k, v)

	}
	fmt.Println("================================================================")
}

func main() {
	test(
		[]string{"John", "Bob", "Jill"},
		[]int{2373772773, 12377383893, 290202992},
	)
	test(
		[]string{"Mike", "Kenny", "Nick"},
		[]int{14252627828, 263573773, 163739309},
	)
	test(
		[]string{"Eric", "Victor", "James"},
		[]int{07366352662, 07363553535, 536635266278},
	)
	test(
		[]string{"Joy", "paul"},
		[]int{07366352662, 07363553535, 536635266278},
	)

	fmt.Println("======USer ID Count================================")
	userIDS := []string{}
	for i := 0; i < 1000; i++ {
		h := md5.New()
		io.WriteString(h, fmt.Sprint(i))
		key := fmt.Sprintf("%x", h.Sum(nil))
		userIDS = append(userIDS, key[:2])
	}

	testCount(userIDS, []string{"00", "ff", "dd"})
	testCount(userIDS, []string{"aa", "12", "32"})
	testCount(userIDS, []string{"bb", "33"})

}
