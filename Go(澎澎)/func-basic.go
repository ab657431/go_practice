package main

import "fmt"

func test(msg string) {
	fmt.Println(msg)
}

func add(n1 int, n2 int) {
	fmt.Println(n1 + n2)
}

func main() {
	var n1 int
	var n2 int
	fmt.Scanln(&n1, &n2)
	add(n1, n2)
}
