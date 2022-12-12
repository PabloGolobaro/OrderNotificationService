package services

import (
	"WhatsappOrderServer/models"
	"WhatsappOrderServer/utils"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
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

func (o *OrderServiceImpl) SaveOrder(order *models.Order) (int, error) {

	order.CreatedAt = time.Now()
	order.UpdatedAt = order.CreatedAt
	orderNumber, err := utils.SetOrderNumber(o.collection, o.ctx)
	order.OrderId = orderNumber
	if err != nil {
		return 0, err
	}
	_, err = o.collection.InsertOne(o.ctx, &order)
	if err != nil {
		return 0, err
	}
	return orderNumber, nil

}
