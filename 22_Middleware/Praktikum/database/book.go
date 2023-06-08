package database

import (
	"go-latihan/config"
	"go-latihan/models"
)

func GetBooks() (interface{}, error) {
	var books []models.Book
	if e := config.DB.Find(&books).Error; e != nil {
		return nil, e
	}
	return books, nil
}

func GetBookById(id int) (interface{}, error) {
	var book models.Book
	if e := config.DB.Find(&book, id).Error; e != nil {
		return nil, e
	}
	return book, nil
}

func Createbook(book *models.Book) (*models.Book, error) {
	if e := config.DB.Create(book).Error; e != nil {
		return nil, e
	}
	return book, nil
}

func DeleteBookById(id int) (interface{}, error) {
	var book models.Book
	if e := config.DB.Delete(&book, id).Error; e != nil {
		return nil, e
	}
	return book, nil
}

func UpdateBookById(id int, updatedBook *models.Book) (*models.Book, error) {
	var book models.Book
	if e := config.DB.First(&book, id).Error; e != nil {
		return nil, e
	}

	book.Judul = updatedBook.Judul
	book.Penulis = updatedBook.Penulis
	book.Penerbit = updatedBook.Penerbit

	if e := config.DB.Save(&book).Error; e != nil {
		return nil, e
	}
	return updatedBook, nil
}
