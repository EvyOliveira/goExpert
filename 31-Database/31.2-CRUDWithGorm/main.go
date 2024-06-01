package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    string `gorm:"primaryKey"`
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
	// products := []Product{
	//	{ID: uuid.New().String(), Name: "Mousepad", Price: 3000.00},
	//	{ID: uuid.New().String(), Name: "Keyboard", Price: 3500.00},
	//	{ID: uuid.New().String(), Name: "Monitor", Price: 4000.00},
	// }
	// db.Create(&products)

	// select one record
	// var product Product
	// db.First(&product, 2)
	// fmt.Println(product)
	// db.First(&product, "name = ?", "Mouse")
	//fmt.Println(product)

	// select all
	// var products []Product
	// db.Find(&products)
	// for _, product := range products {
	//   fmt.Println(product)
	// }

	// get two registers at the same time
	// var products []Product
	// db.Limit(2).Find(&products)
	// for _, product := range products {
	//	fmt.Println(product)
	// }

	// get two registers at the same time using pagination
	var products []Product
	db.Limit(2).Offset(2).Find(&products)
	for _, product := range products {
		fmt.Println(product)
	}
}
