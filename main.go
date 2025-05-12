package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	controlsInGo()
}

func controlsInGo() {
	num := 12

	switch {
	case num < 10:
		fmt.Println("Number is less than 10")
	case num < 20:
		fmt.Println("Number is less than 20")
	}
}
