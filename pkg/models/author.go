package models

import (
	"github.com/jinzhu/gorm"
	"github.com/maratb3k/go-bookstore/pkg/config"
)

type Author struct {
	gorm.Model
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	MiddleName string `json:"middle_name"`
	books      []Book `gorm:"foreignKey:AuthorId"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Author{})
}

func (a *Author) CreateAuthor() *Author {
	db.NewRecord(a)
	db.Create(&a)
	return a
}

func GetAllAuthors() []Author {
	var Authors []Author
	db.Find(&Authors)
	return Authors
}

func GetAuthorById(Id int64) (*Author, *gorm.DB) {
	var getAuthor Author
	db := db.Where("Id = ?", Id).Find(&getAuthor)
	return &getAuthor, db
}

func DeleteAuthorById(ID int64) Author {
	var author Author
	db.Where("Id = ?", ID).Delete(&author)
	return author
}
