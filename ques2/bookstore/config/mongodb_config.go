package config

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	*mongo.Database
}

var mongoDbConn *MongoDB

func ConfigureMongoDB() error {
	bookStoreDBProps := map[string]string{
		"database": "bookstore",
		"username": "kaushal",
		"password": "password_mongo",
		"host":     "cluster0.mevyk.mongodb.net",
	}

	connURI := getMongoDBConnString(bookStoreDBProps)
	clientOptions := options.Client().ApplyURI(connURI)
	ctx := context.Background()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return err
	}

	db := client.Database(bookStoreDBProps["database"])
	mongoDbConn = &MongoDB{db}
	return nil
}

func GetMongoDBClient() *MongoDB {
	return mongoDbConn
}

func getMongoDBConnString(dbProps map[string]string) string {
	return fmt.Sprintf("mongodb+srv://%s:%s@%s/%s?retryWrites=true&w=majority",
		// return fmt.Sprintf("mongodb+srv://%s:${encodeURIComponent(%s)}@%s/%s?retryWrites=true&w=majority",
		dbProps["username"],
		dbProps["password"],
		dbProps["host"],
		dbProps["database"])
}
