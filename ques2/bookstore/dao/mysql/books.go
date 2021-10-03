package mysql

import (
	"bookstore/config"
	"bookstore/models"
	"log"
)

type MySQLBookStoreDAO struct {
}

func (store *MySQLBookStoreDAO) Database() string {
	return "bookstore"
}

func (store *MySQLBookStoreDAO) CreateBook(book models.Book) error {
	db := config.GetMysqlClient()
	result := db.Create(&book)
	if result.Error != nil {
		return result.Error
	}
	rowsAffected := result.RowsAffected
	log.Printf("Rows affected %d", rowsAffected)

	return nil
}

func (store *MySQLBookStoreDAO) GetBookByTitle(title string) ([]models.Book, error) {
	var err error
	db := config.GetMysqlClient()
	var books []models.Book
	if err = db.Where("title = ?", title).Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}
