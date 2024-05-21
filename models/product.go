package models

import "gorm.io/gorm"

type Product struct{
    gorm.Model
    ID      int64 `gorm:"primaryKey" json:"id"`
    Name    string `json:"name"`
    Price   float64 `json:"price"`
    Stock   int `json:"stock"`
}
