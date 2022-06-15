package main

import "fmt"

func add(x *int) {
	*x += 10
	fmt.Println(*x)
}
func main() {
	var x int = 10
	var x_ptr *int = &x
	add(x_ptr)
	fmt.Println(x)
}
