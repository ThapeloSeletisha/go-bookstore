package models

import (
	"github.com/ThapeloSeletisha/go-bookstore/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model 
	Name string `gorm:"json:name"` 
	Author string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

// Enters record of new book into database
func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

// Obtains all records of books in database
func GetAllBooks() []Book {
	var books []Book // akhil books are called Books
	db.Find(&books)
	return books
}

// Obtains record of book with the given ID
func GetBookById(ID int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", ID).Find(&getBook)
	return &getBook, db
}

// Deletes record of book with the given ID
func DeleteBook(ID int64) Book {
	var book Book
	db.Where("ID=?", ID).Delete(&book)
	return book
}