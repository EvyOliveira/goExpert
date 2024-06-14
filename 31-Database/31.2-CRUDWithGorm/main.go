package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    string `gorm:"primaryKey"`
	Name  string
	Price float64
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
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
	// var products []Product
	// db.Limit(2).Offset(2).Find(&products)
	// for _, product := range products {
	//	fmt.Println(product)
	//}

	// select records using where
	// var products []Product
	// db.Where("price > ?", 100).Find(&products)
	// for _, product := range products {
	//	fmt.Println(product)
	//}

	// select records using like
	// var products []Product
	// db.Where("name LIKE ?", "%book%").Find(&products)
	// for _, product := range products {
	//	fmt.Println(product)
	//}

	// modify first record name and delete the secound record
	// var p Product
	// db.First(&p, 1)
	// p.Name = "New Mouse"
	// db.Save(&p)

	// var p2 Product
	// db.First(&p2, 1)
	// fmt.Println(p2.Name)
	// db.Delete(&p2)
}
