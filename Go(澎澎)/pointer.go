package main

import "fmt"

func main() {
	//1.建立變數存放資料
	var x int = 5
	fmt.Println("原來的資料:", x)
	//2.取得記憶體位置:&變數名稱
	fmt.Println("資料的記憶體位置:", &x)
	//3.存放到指標變數。指標變數資料型態:*資料型態
	var xtr *int = &x
	fmt.Println("資料的記憶體位置:", xtr)
	//4.反解指標變數:*指標變數名稱
	fmt.Println("反解指標xPtr回原本的資料", *xtr)
}
