package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primaryKey;autoIncrement"`
	Name  string
	Price float64
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{})

	// create a record
	// db.Create(&Product{
	// 	Name:  "Notebook",
	// 	Price: 2000.00,
	// })

	// create batch register
	products := []Product{
		{Name: "Mousepad", Price: 3000.00},
		{Name: "Keyboard", Price: 3500.00},
		{Name: "Monitor", Price: 4000.00},
	}
	db.Create(&products)
}
