package main

import (
	"bookstore/config"
	"bookstore/dao/interfaces"
	"bookstore/models"
	"flag"
	"log"

	mongodbdao "bookstore/dao/mongodb"
	mysqldao "bookstore/dao/mysql"
)

var dbEngine string

func init() {
	// default database is mysql
	databasePtr := flag.String("database", "mysql", "database for persistent storage")
	flag.Parse()
	dbEngine = *databasePtr
}

func configureDatabase(dbEngine string) interfaces.BookStoreDAO {
	var bookStore interfaces.BookStoreDAO
	if dbEngine == "mysql" {
		err := config.ConfigureMySql()
		if err != nil {
			log.Printf("Error occured while setting up MySQL connection %s", err.Error())
			panic(err.Error())
		}
		db := config.GetMysqlClient()
		db.AutoMigrate(&models.Book{})
		log.Println("connected to mysql successfully")
		bookStore = &mysqldao.MySQLBookStoreDAO{}
	} else if dbEngine == "mongodb" {
		err := config.ConfigureMongoDB()
		if err != nil {
			log.Printf("Error occured while setting up mongodb connection %s", err.Error())
			panic(err.Error())
		}
		log.Println("connected to mongodb successfully")
		bookStore = &mongodbdao.MongoDBBookStoreDAO{}
	} else {
		log.Fatal("ERROR: Unknown database engine")
	}
	return bookStore
}

/*
	--- commands to run ---
	go run main.go
	go run main.go -database mongodb
*/
func main() {
	bs := configureDatabase(dbEngine)

	book1 := models.Book{
		Isbn:   "isbn1",
		Title:  "DDIA",
		Author: "Martin Kleppmann",
		Price:  1600,
	}
	err := bs.CreateBook(book1)
	if err != nil {
		log.Printf("Error while creating book %s", err.Error())
	}
	log.Println("Record added to database successfully")

	log.Println("Searching booko with title DDIA")
	books, err := bs.GetBookByTitle("DDIA")
	if err != nil {
		log.Printf("Error while searching the book title: %s", err.Error())
	}
	log.Println(books)
}
