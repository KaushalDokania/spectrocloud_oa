package interfaces

import "bookstore/models"

type BookStoreDAO interface {
	CreateBook(book models.Book) error
	GetBookByTitle(title string) ([]models.Book, error)
}
