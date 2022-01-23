package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	username := []string{"jack", "mack", "sack"}
	password := []int{1234, 123, 345}
	var indexNum int

	for index, value := range username {

		if value == os.Args[1] {
			indexNum = index
			break
		}

	}

	if os.Args[1] == username[indexNum] {

		if os.Args[2] == strconv.Itoa(password[indexNum]) {

			fmt.Printf(" Access granted to user %v ", os.Args[1])
		} else {
			fmt.Printf("Incorrect password for user %v", os.Args[1])
		}
	} else {

		fmt.Printf(" User %v not present", os.Args[1])
	}
}
