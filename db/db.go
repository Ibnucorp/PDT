package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB()(*gorm.DB,error){
    conn := "root:@tcp(127.0.0.1:3306)/ecommerce?charset=utf8mb4&parseTime=True&loc=UTC"
    db,err := gorm.Open(mysql.Open(conn), &gorm.Config{})

    if (err != nil){
        return nil,err
    }

    DB = db
    return DB,nil
}