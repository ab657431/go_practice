package main

import "fmt"

var money int

func main() {
	fmt.Scanln(&money)
	if money < 100000 {
		fmt.Print("ok")
	} else {
		fmt.Print("too much")
	}
}
