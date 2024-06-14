package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Product struct {
	ID           string `gorm:"primaryKey"`
	Name         string
	Price        float64
	Categories   []Category `gorm:"many2many:products_categories;"`
	SerialNumber SerialNumber
	gorm.Model
}

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product `gorm:"many2many:products_categories;"`
}

type SerialNumber struct {
	ID        int `gorm:"primaryKey"`
	Number    string
	ProductID int
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

	category := Category{Name: "Cozinha"}
	db.Create(&category)

	category2 := Category{Name: "Eletrônicos"}
	db.Create(&category2)

	// create category
	// category := Category{Name: "Eletronicos"}
	// db.Create(&category)

	// create product
	db.Create(&Product{
		Name:       "Panela",
		Price:      99.00,
		Categories: []Category{category, category2},
	})

	// create serial number
	db.Create(&SerialNumber{
		Number:    "123456",
		ProductID: 1,
	})

	var products []Product
	db.Preload("Category").Preload("SerialNumber").Find(&products)
	for _, product := range products {
		fmt.Println(product.Name, product.SerialNumber.Number)
	}

	var categories []Category
	err = db.Model(&Category{}).Preload("Products").Find(&categories).Error
	if err != nil {
		panic(err)
	}
	for _, category := range categories {
		fmt.Println(category.Name, ":")
		for _, product := range category.Products {
			println("-", product.Name)
		}
	}

	// select * from products where id = 1 FOR UPDATE
	// uma forma de se trabalhar com lock pessimista
	tx := db.Begin()
	var c Category
	err = tx.Debug().Clauses(clause.Locking{Strength: "UPDATE"}).First(&c, 1).Error
	if err != nil {
		panic(err)
	}
	c.Name = "Eletronicos"
	tx.Debug().Save(&c)
	tx.Commit()
}
