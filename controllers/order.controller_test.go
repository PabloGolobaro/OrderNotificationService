package controllers

import (
	"WhatsappOrderServer/config"
	"WhatsappOrderServer/models"
	"WhatsappOrderServer/services"
	bytes2 "bytes"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/magiconair/properties/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestOrderController_SaveOrder(t *testing.T) {
	type fields struct {
		orderService services.OrderService
		ctx          context.Context
	}
	test_fields := fields{}
	test_fields.ctx = context.TODO()
	//connect to MongoDB
	config, _ := config.LoadConfig("C:\\Users\\Professional\\GolandProjects\\WhatsappOrderServer")
	clientOptions := options.Client().ApplyURI(config.DBUri)
	mongoclient, err := mongo.Connect(test_fields.ctx, clientOptions)
	if err != nil {
		panic(err)
	}
	orderCollection := mongoclient.Database("golang_mongodb").Collection("orders")
	test_fields.orderService = services.NewOrderServiceImpl(orderCollection, test_fields.ctx)
	oc := &OrderController{orderService: test_fields.orderService, ctx: test_fields.ctx}
	//server
	r := gin.Default()
	r.POST("/api/order", oc.SaveOrder)

	tests := []struct {
		name string
	}{{name: "first"}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			order := &models.Order{
				Products:  []models.Product{{Id: 1, Category: "Food", Img: "123.png", Title: "Shashlik", Price: "120", Count: 3}, {Id: 2, Category: "Food", Img: "333.png", Title: "Sauce", Price: "30", Count: 1}},
				Data:      models.OrderData{Adress: "Orel, sm. Per 10", Mobile: "89991002345671", Time: time.Now().String()},
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}
			bytes, _ := json.Marshal(order)
			reader := bytes2.NewReader(bytes)
			w := httptest.NewRecorder()
			request, _ := http.NewRequest(http.MethodPost, "/api/order", reader)
			r.ServeHTTP(w, request)
			assert.Equal(t, http.StatusCreated, w.Code, w.Body.String())
		})
	}
}
