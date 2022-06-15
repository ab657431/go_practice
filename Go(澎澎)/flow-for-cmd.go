package main

import "fmt"

func main() {
	/* var x int = 0
	for x = 0; x < 5; x++ {
		if x == 3 {
			continue
		}
		fmt.Println(x)

	}*/
	var result int = 0
	for true {
		var x int
		fmt.Scanln(&x)
		if x == 0 {
			break
		} else {
			result += x
		}
		fmt.Println(result)
	}
}
