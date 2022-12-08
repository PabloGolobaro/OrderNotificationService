package main

import (
	"WhatsappOrderServer/config"
	"WhatsappOrderServer/controllers"
	"WhatsappOrderServer/routes"
	"WhatsappOrderServer/services"
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"net/http"
)

var (
	server      *gin.Engine
	ctx         context.Context
	mongoclient *mongo.Client
	logger      log.Logger

	orderCollection      *mongo.Collection
	orderService         services.OrderService
	OrderController      controllers.OrderController
	OrderRouteController routes.OrderRouteController
)

func main() {

	config, err := config.LoadConfig(".")

	if err != nil {
		logger.Fatal("Could not load config", err)
	}

	defer mongoclient.Disconnect(ctx)

	apiGroup := server.Group("/api")
	apiGroup.GET("/healthchecker", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Hello from PabloGolobar"})
	})

	OrderRouteController.AuthRoute(apiGroup)
	logger.Fatal(server.Run(":" + config.Port))

}

func init() {
	logger.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	config, err := config.LoadConfig(".")
	if err != nil {
		logger.Fatal("Could not load environment variables", err)
	}

	ctx = context.TODO()
	clientOptions := options.Client().ApplyURI(config.DBUri)
	mongoclient, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(err)
	}

	if err := mongoclient.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	logger.Println("MongoDB successfully connected...")

	// collections
	orderCollection = mongoclient.Database("golang_mongodb").Collection("orders")
	orderService = services.NewOrderServiceImpl(orderCollection, ctx)
	OrderController = controllers.NewOrderController(orderService, ctx)
	OrderRouteController = routes.NewOrderRouteController(OrderController)

	server = gin.Default()
}
