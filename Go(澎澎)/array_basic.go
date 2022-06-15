package main

import "fmt"

func main() {
	//整數陣列
	/*
		var numbers [3]int
		numbers[0] = 1
		numbers[1] = 2
		fmt.Println(numbers)
		//字串陣列
		var names [3]string = [3]string{"bossn", "bryant", "kobe"}
		fmt.Println(names)
		//取得陣列長度
		fmt.Println(len(names))
		//逐一取得陣列中的資料
	*/
	var grades [3]int
	var grade int
	var sum int
	fmt.Println("請輸入數字,系統將為您計算總合")
	for grade = 0; grade < len(grades); grade++ {
		fmt.Scanln(&grades[grade])
	}
	for grade = 0; grade < len(grades); grade++ {
		sum += grades[grade]
	}
	fmt.Println("平均分數為:", sum/len(grades))
}
