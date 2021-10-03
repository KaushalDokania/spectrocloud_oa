package mongodb

import (
	"bookstore/config"
	"bookstore/models"
	"context"
	"log"
	"math/rand"
	"time"

	"go.mongodb.org/mongo-driver/bson"
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
	var err error

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

func (store *MongoDBBookStoreDAO) GetBookByTitle(title string) ([]models.Book, error) {
	var err error
	db := config.GetMongoDBClient()
	collection := db.Collection(store.Collection())
	searchQuery := bson.M{"title": title}
	cursor, err := collection.Find(context.Background(), searchQuery)

	if err != nil {
		return nil, err
	}
	var books []models.Book
	if err = cursor.All(context.Background(), &books); err != nil {
		return nil, err
	}
	return books, err
}
