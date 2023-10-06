package models

import (
	"github.com/Kawar1mi/go-bookstore/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"name" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.Create(&b)
	return b
}

func GetBooks() (Books []Book) {
	db.Find(&Books)
	return
}

func GetBookById(id int64) (*Book, *gorm.DB) {
	var book Book
	db := db.Where("ID=?", id).First(&book)
	return &book, db
}

func DeleteBook(id int64) (book Book) {
	db.Where("ID = ?", id).Delete(&book)
	return
}
