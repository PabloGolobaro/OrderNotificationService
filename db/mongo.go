package db

import (
	"WhatsappOrderServer/config"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
)

func NewMongoDBConnection(config config.Config, logger *log.Logger) (*mongo.Client, error) {
	ctx := context.TODO()
	clientOptions := options.Client().ApplyURI(config.DBUri)
	mongoclient, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(err)
	}

	if err := mongoclient.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	logger.Println("MongoDB successfully connected...")
	return mongoclient, nil
}
