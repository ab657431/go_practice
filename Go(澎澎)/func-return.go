package main

import "fmt"

func show(msg string) string {
	fmt.Println(msg)
	return msg
}

func add(n1 int, n2 int) int {
	fmt.Println(n1 + n2)
	return n1 + n2
}

func main() {

	var x string = show("你好")
	fmt.Println(x)
	var y int = add(3, 4)
	fmt.Println(y)
}
