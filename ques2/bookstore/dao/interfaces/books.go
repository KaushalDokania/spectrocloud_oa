package interfaces

import "bookstore/models"

type BookStoreDAO interface {
	CreateBook(book models.Book) error
}
