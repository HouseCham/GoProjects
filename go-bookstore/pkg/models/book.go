package models

import (
	"github.com/HouseCham/go-bookstore/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name string `gorm:""json:"name"`
	Author string `json:"author"`
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

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB){
	var bookObtain Book
	db := db.Where("Id=?", Id).Find(&bookObtain)
	return &bookObtain, db
}

func DeleteBook(Id int64) Book {
	var book Book
	db.Where("id = ?", Id).Delete(&book)
	return book
}