package models

import (
	"github.com/jinzhu/gorm"
	"github.com/maratb3k/go-bookstore/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name string `gorm:"" json:"name"`
	AuthorId      uint `json:"author"`
	Publication string `json:"publication"`
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
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("Id = ?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBookById(ID int64) Book {
	var book Book
	db.Where("Id = ?", ID).Delete(&book)
	return book
}
