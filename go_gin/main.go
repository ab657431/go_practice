package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

//handler(GET方法)
func LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

//handler(POST方法)
func LoginAuth(c *gin.Context) {

	//先宣告使用者資料的型態
	var (
		username string
		password string
	)
	//確認使用者是否有輸入使用者名稱
	//GetPostForm返回兩個引數，第一個是引數值，第二個引數是引數是否存在的bool值
	if in, isExist := c.GetPostForm("username"); isExist && in != "" {
		username = in
	} else {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"error": errors.New("必須輸入使用者名稱"),
		})
		return
	}
	if in, isExist := c.GetPostForm("password"); isExist && in != "" {
		password = in
	} else {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"error": errors.New("必須輸入密碼名稱"),
		})
		return
	}
	//Auth函式需傳入使用者名稱、密碼，若使用者名稱及密碼接比對成功,return nil
	//比對失敗,回傳	errors.New("user is not exist")
	if err := Auth(username, password); err == nil {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"success": "登入成功",
		})
		return
	} else {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{
			"error": err,
		})
		return
	}
}

func main() {
	server := gin.Default()
	server.LoadHTMLGlob("template/html/*")
	//靜態資源的讀取
	server.Static("/assets", "./template/assets")
	//route及handler
	server.GET("/login", LoginPage)
	server.POST("/login", LoginAuth)
	server.Run(":7777")
}
