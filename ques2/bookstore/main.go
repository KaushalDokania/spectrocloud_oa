package main

import (
	"bookstore/config"
	"bookstore/models"
	"flag"
	"log"
)

var dbEngine string

func init() {
	// default database is mysql
	databasePtr := flag.String("database", "mysql", "database for persistent storage")
	flag.Parse()
	dbEngine = *databasePtr
}

func configureDatabase(dbEngine string) {
	if dbEngine == "mysql" {
		err := config.ConfigureMySql()
		if err != nil {
			log.Printf("Error occured while setting up MySQL connection %s", err.Error())
			panic(err.Error())
		}
		db := config.GetMysqlClient()
		db.AutoMigrate(&models.Book{})
		log.Println("connected to mysql successfully")
	} else if dbEngine == "mongodb" {
		err := config.ConfigureMongoDB()
		if err != nil {
			log.Printf("Error occured while setting up mongodb connection %s", err.Error())
			panic(err.Error())
		}
		log.Println("connected to mongodb successfully")
	} else {
		log.Fatal("ERROR: Unknown database engine")
	}
}

/*
	--- commands to run ---
	go run main.go
	go run main.go -database mongodb
*/
func main() {
	log.Println("Hello World !!")
	configureDatabase(dbEngine)
}
