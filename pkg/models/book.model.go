package models

import (
	"bookstore-go/pkg/config"

	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDb()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.Create(&b)
	return b
}

// get all books
func GetAllBooks() []*Book {
	var books []*Book
	db.Find(&books)
	return books
}

// get book by id
func GetBookById(id int64) (*Book, error) {
	var book *Book
	err := db.Where("ID=?", id).Find(&book).Error
	return book, err
}

// delete book by id
func DeleteBookById(id int64) error {
	return db.Where("ID=?", id).Delete(&Book{}).Error
}

// update books by id
func (b *Book) UpdateBook(id int64) *Book {
	db.Where("ID=?", id).Updates(&b)
	return b
}
