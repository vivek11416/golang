package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Printf("%#v\n", os.Args)

	fmt.Println("Path : ", len(os.Args))
}
