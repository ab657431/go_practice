package main //參考https://ithelp.ithome.com.tw/articles/10233981

import (
	"html/template"
	"log"
	"net/http"
)

//建立結構
type IndexData struct {
	Title   string
	Content string
}

//
func test(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./index.html"))
	data := new(IndexData)
	data.Title = "首頁"
	data.Content = "我的第一個首頁"
	tmpl.Execute(w, data)
}

func main() {
	http.HandleFunc("/", test)
	http.HandleFunc("/index", test)
	err := http.ListenAndServe(":8877", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
