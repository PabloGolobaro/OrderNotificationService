package services

import (
	"WhatsappOrderServer/config"
	"WhatsappOrderServer/models"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"testing"
	"time"
)

func TestOrderServiceImpl_SaveOrder(t *testing.T) {
	type fields struct {
		collection *mongo.Collection
		ctx        context.Context
	}
	test_fields := fields{}
	test_fields.ctx = context.TODO()
	//connect to MongoDB
	config, err := config.LoadConfig("C:\\Users\\Professional\\GolandProjects\\WhatsappOrderServer")
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}
	clientOptions := options.Client().ApplyURI(config.DBUri)
	mongoclient, err := mongo.Connect(test_fields.ctx, clientOptions)
	if err != nil {
		panic(err)
	}
	if err := mongoclient.Ping(test_fields.ctx, readpref.Primary()); err != nil {
		panic(err)
	}
	log.Println("MongoDB successfully connected...")
	orderCollection := mongoclient.Database("golang_mongodb").Collection("orders")
	test_fields.collection = orderCollection
	//test_struct
	type args struct {
		order *models.Order
	}

	arguments := args{order: &models.Order{
		Products:  []models.Product{{Id: 1, Category: "Food", Img: "123.png", Title: "Shashlik", Price: "120", Count: 3}, {Id: 2, Category: "Food", Img: "333.png", Title: "Sauce", Price: "30", Count: 1}},
		Data:      models.OrderData{Adress: "Orel, sm. Per 10", Mobile: "89991002345671", Time: time.Now().String()},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{

		{name: "First", fields: test_fields, args: arguments, want: "some", wantErr: false},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &OrderServiceImpl{
				collection: tt.fields.collection,
				ctx:        tt.fields.ctx,
			}
			got, err := o.SaveOrder(tt.args.order)
			if (err != nil) != tt.wantErr {
				t.Errorf("SaveOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SaveOrder() got = %v, want %v", got, tt.want)
			}
		})
	}
}
