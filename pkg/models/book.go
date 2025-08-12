package models

import (
	"github.com/Kushagra-KK/BookStore/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:""json:"name"`
	Author      string `json:"author"`
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
	db := db.Where("Id=?", Id).Find(&getBook)
	return &getBook, db
}
func DeleteBook(Id int64) Book {
	var book Book
	db.Where("Id=?", Id).Delete(book)
	return book
}
