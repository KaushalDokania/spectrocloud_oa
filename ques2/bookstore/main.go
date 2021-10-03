package main

import (
	"bookstore/config"
	"bookstore/models"
	"log"
)

func main() {
	log.Println("Hello World !!")

	err := config.ConfigureMySql()
	if err != nil {
		log.Printf("Error occured while setting up MySQL connection %s", err.Error())
		panic(err.Error())
	}
	db := config.GetMysqlClient()
	db.AutoMigrate(&models.Book{})
	log.Println("connected to mysql successfully")
}
