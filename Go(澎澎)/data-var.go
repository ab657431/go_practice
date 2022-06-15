package main //生成執行檔需使用main封包
import "fmt" //載入內建封包,用於基本輸出輸入
func main() { //建立main函式,程式的進入點
	//執行程式時,從這邊開始執行
	//fmt.Println('A') //字符
	//變數的使用
	var x int = 4 //宣告變數的資料型態
	fmt.Println(x)
	var y rune = 'c'
	fmt.Println(y)
}
