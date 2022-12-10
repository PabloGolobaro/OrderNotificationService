package services

import (
	"WhatsappOrderServer/models"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"strings"
	"time"
)

type OrderServiceImpl struct {
	collection *mongo.Collection
	ctx        context.Context
}

func (o *OrderServiceImpl) GetAllOrder() ([]*models.DBResponse, error) {
	//TODO implement me
	panic("implement me")
}

func NewOrderServiceImpl(collection *mongo.Collection, ctx context.Context) OrderService {
	return &OrderServiceImpl{collection: collection, ctx: ctx}
}

func (o *OrderServiceImpl) SaveOrder(order *models.Order) (string, error) {
	order.CreatedAt = time.Now()
	order.UpdatedAt = order.CreatedAt
	res, err := o.collection.InsertOne(o.ctx, &order)
	if err != nil {
		return "", err
	}
	s := res.InsertedID.(primitive.ObjectID).String()
	split := strings.Split(s, "\"")
	return split[1], nil

}
