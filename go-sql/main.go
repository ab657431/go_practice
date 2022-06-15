package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	USERNAME = "root"
	PASSWORD = "ab563131"
	NETWORK  = "tcp"
	SERVER   = "127.0.0.1"
	PORT     = 3306
	DATABASE = "demo"
)

//創建table,把sql語句指定給變數sql
func CreateTable(db *sql.DB) error {
	sql := `CREATE TABLE IF NOT EXISTS users(
	id INT(4) PRIMARY KEY AUTO_INCREMENT NOT NULL,
        username VARCHAR(64),
        password VARCHAR(64)
	); `
	//若err不為nil表示有錯誤發生,這時return錯誤訊息
	if _, err := db.Exec(sql); err != nil {
		fmt.Println("建立 Table 發生錯誤:", err)
		return err
	}
	//若無錯誤產生,print建立成功,並回傳nil
	fmt.Println("建立 Table 成功！")
	return nil
}

//寫入資料表
func InsertUser(DB *sql.DB, username, password string) error {
	_, err := DB.Exec("insert INTO users(username,password) values(?,?)", username, password)
	if err != nil {
		fmt.Printf("建立使用者失敗，原因是：%v", err)
		return err
	}
	fmt.Println("建立使用者成功！")
	return nil
}

type User struct {
	ID       string
	Username string
	Password string
}

//查詢資料表
func QueryUser(db *sql.DB, username string) {
	user := new(User)
	row := db.QueryRow("select * from users where username=?", username)
	if err := row.Scan(&user.ID, &user.Username, &user.Password); err != nil {
		fmt.Printf("映射使用者失敗，原因為：%v\n", err)
		return
	}
	fmt.Println("查詢使用者成功", *user)
}

func main() {
	conn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	db, err := sql.Open("mysql", conn)
	if err != nil {
		fmt.Println("開啟 MySQL時發生錯誤,原因是:", err)
		return
	}
	if err := db.Ping(); err != nil {
		fmt.Println("資料庫連線錯誤，原因為：", err.Error())
		return
	}
	defer db.Close()
	CreateTable(db)
	InsertUser(db, "ab657431", "ab563131")
	QueryUser(db, "ab657431")
}
