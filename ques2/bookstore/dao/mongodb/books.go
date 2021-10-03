package mongodb

import (
	"bookstore/config"
	"bookstore/models"
	"context"
	"log"
	"math/rand"
	"time"
)

type MongoDBBookStoreDAO struct {
}

func (store *MongoDBBookStoreDAO) Database() string {
	return "bookstore"
}

func (store *MongoDBBookStoreDAO) Collection() string {
	return "books"
}

func (store *MongoDBBookStoreDAO) CreateBook(book models.Book) error {
	rand.Seed(time.Now().UnixNano())
	id := 1000 + rand.Intn(1000)
	book.ID = uint64(id)
	db := config.GetMongoDBClient()
	collection := db.Collection(store.Collection())
	result, err := collection.InsertOne(context.TODO(), book)
	if err != nil {
		return err
	}
	insertedId := result.InsertedID
	log.Printf("Created Book with ID: %d", insertedId)
	return nil
}
