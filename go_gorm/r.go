package gogorm

import "fmt"

func te() {
	// var p *int  // nil

	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(p)  // 0xc0000b4008
	fmt.Println(*p) // 透過 pointer 來讀取到 i
	*p = 21         // 透過 pointer 來設定 i 的值
	fmt.Println(i)  // 21

	p = &j         // 將 p 的值設為 j 的 pointer
	*p = *p / 37   // 透過 pointer 來除 j
	fmt.Println(j) // 73
}
