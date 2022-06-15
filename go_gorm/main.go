package main

//go get -u gorm.io/gorm
//go get -u gorm.io/driver/mysql

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//連線資料
const (
	USERNAME = "root"
	PASSWORD = "ab563131"
	NETWORK  = "tcp"
	SERVER   = "127.0.0.1"
	PORT     = 3306
	DATABASE = "demo"
)

//需要migrate的table資料,table取名方式為
//Bosn+s
type User struct {
	ID       int64  `json:"id" gorm:"primary_key;auto_increase'"`
	Username string `json:"username"`
	Password string `json:""`
}

func CreateUser(db *gorm.DB, user *User) error {
	return db.Create(user).Error
}

func FindUser(db *gorm.DB, id int64) (*User, error) {
	user := new(User)
	user.ID = id
	err := db.First(&user).Error
	return user, err
}

func main() {
	//建立連線
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("使用 gorm 連線 DB 發生錯誤，原因為 " + err.Error())
	}
	//建立table
	if err := db.AutoMigrate(new(User)); err != nil {
		panic("資料庫 Migrate 失敗，原因為 " + err.Error())
	}
	//建立一個要寫入的資料後migrate
	user := &User{
		Username: "test",
		Password: "test",
	}
	//創建使用者
	if err := CreateUser(db, user); err != nil {
		panic("資料庫 Migrate 失敗，原因為 " + err.Error())
	}
	//查詢使用者
	if user, err := FindUser(db, 9); err == nil {
		log.Println("查詢到的user為", user)
	} else {
		panic("查詢 user 失敗,原因為" + err.Error())
	}

}
