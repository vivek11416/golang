//path package helps to work with URLs
package main

import (
	"fmt"
	"path"
)

func main2() {

	dir, file := path.Split("css/main.css")

	fmt.Print("dir : ", dir)
	fmt.Print("file : ", file)

	speed := 100
	force := 2.5

	speed = int(float64(speed) * force)

	fmt.Println(speed)

}
