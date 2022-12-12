package utils

import (
	"WhatsappOrderServer/models"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func SetOrderNumber(collection *mongo.Collection, ctx context.Context) (int, error) {
	var order models.DBResponse
	findOptions := options.Find().SetSort(bson.D{{"order_id", -1}}).SetLimit(1)
	cursor, err := collection.Find(ctx, bson.D{}, findOptions)
	if err != nil {
		if err == mongo.ErrNilCursor {
			log.Println(err)
			return 1, nil
		}
		return 0, err
	}
	for cursor.Next(ctx) {
		err := cursor.Decode(&order)
		if err != nil {
			return 0, err
		}
	}
	return order.OrderId + 1, nil
}
