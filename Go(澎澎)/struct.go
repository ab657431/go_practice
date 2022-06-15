package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	var p1 Person = Person{"恩恩", 26}
	fmt.Println(p1.name, p1.age)
	var p2 Person = Person{name: "哭吉", age: 1}
	fmt.Println(p2.name, p2.age)
}
