package main

import (
	"fmt"
	"os"
)

func main3() {

	strVal := "Hello"

	fmt.Println(len(strVal)) // len counts the number of bytes
	fmt.Printf("%#v\n", os.Args)

	fmt.Println("Path : ", len(os.Args)) //
}
