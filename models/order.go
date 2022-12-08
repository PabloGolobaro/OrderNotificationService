package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Product struct {
	Id       int    ` bson:"id" json:"id" `
	Category string ` bson:"category" json:"category"`
	Img      string ` bson:"img" json:"img"`
	Title    string ` bson:"title" json:"title"`
	Price    string ` bson:"price" json:"price"`
	Count    int    ` bson:"count" json:"count"`
}

type OrderData struct {
	Adress string `bson:"adress" json:"adress"`
	Mobile string `bson:"mobile" json:"mobile"`
	Time   string `bson:"time" json:"time"`
}

type Order struct {
	Products  []Product
	Data      OrderData
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

type DBResponse struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Products  []Product
	Data      OrderData
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}
