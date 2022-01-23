package main

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func main4() {

	//fmt.Println(utf8.RuneCountInString(os.Args[1]))
	fmt.Println(strings.ToUpper(os.Args[1]) + strings.Repeat("!", utf8.RuneCountInString(os.Args[1])))
}
