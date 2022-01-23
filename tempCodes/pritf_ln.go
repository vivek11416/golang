package main

import "fmt"

func main6() {
	//%d = Integer
	//%v = verb
	//%f = float
	//%s = string
	//%t = Boolean
	name := "vivek"
	age := 24.254
	name = "Rahul"
	fmt.Printf("%v is my name , and %.2f is my age . Imagine i am %[2]v !!!!!!!", name, age)
}
