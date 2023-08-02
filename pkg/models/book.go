package models

import (
	"github.com/jinzhu/gorm"
	"github.com/programmingfire/bookstore/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Title       string `gorm:""json:"title"`
	Author      string `gorm:""json:"author"`
	Publication string `gorm:""json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var books []Book
	db.Find(&books)
	return books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var book Book
	db.Where("id =?", Id).First(&book)
	return &book, db
}

func DeleteBook(Id int64) Book {
	var book Book
	db.Where("id =?", Id).Delete(&book)
	return book
}
