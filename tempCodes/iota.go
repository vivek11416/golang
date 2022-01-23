package main

import "fmt"

func main5() {
	//iota is a built -on constant generator which generates ever increasing numbers
	const (
		sun = iota + 1
		_   // blank identifier
		mon
		tue
		wed
		thu
		fri
		sat
	)

	fmt.Println(sun, mon, tue, wed, thu, fri, sat)
}
