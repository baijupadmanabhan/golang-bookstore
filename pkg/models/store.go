package models

import (
	"github.com/baijupadmanabhan/golang-bookstore/pkg/config"
	"github.com/jinzhu/gorm"
)

var dbConn *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	dbConn = config.GetDBConn()
	dbConn.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	dbConn.NewRecord(b)
	dbConn.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Store []Book
	dbConn.Find(&Store)
	return Store
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := dbConn.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(ID int64) Book {
	var book Book
	dbConn.Where("ID=?", ID).Delete(book)
	return book
}
