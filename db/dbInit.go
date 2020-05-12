package db

import (
	model "api/model"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //_ 只是调用init函数
)

var DB *gorm.DB

//
func DbInit() {
	var err error
	DB, err = gorm.Open("mysql", "haige:password@(myfirstmysql.cazjqlq1ycv3.us-east-2.rds.amazonaws.com:3306)/mydb?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("error catched")
	}

	//snake game logic
	DB.AutoMigrate(&model.GoUser{})
}

func CloseDb() {
	DB.Close()
}
