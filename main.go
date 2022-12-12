package main

import (
	"WhatsappOrderServer/config"
	"WhatsappOrderServer/controllers"
	"WhatsappOrderServer/db"
	"WhatsappOrderServer/handlers"
	"WhatsappOrderServer/routes"
	"WhatsappOrderServer/services"
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	tele "gopkg.in/telebot.v3"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	server      *gin.Engine
	ctx         context.Context
	mongoclient *mongo.Client
	logger      *log.Logger

	bot *tele.Bot

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
	go func() {
		bot.Start()
		logger.Println("Started bot!")
	}()
	OrderRouteController.AuthRoute(apiGroup)
	logger.Fatal(server.Run(":" + config.Port))

}

func init() {
	logger = log.New(os.Stdout, "server", log.Ldate|log.Ltime|log.Lshortfile)
	config, err := config.LoadConfig(".")
	if err != nil {
		logger.Fatal("Could not load environment variables", err)
	}
	mongoclient, err = db.NewMongoDBConnection(config, logger)
	if err != nil {
		return
	}

	bot = NewBot(config)

	// collections
	orderCollection = mongoclient.Database("golang_mongodb").Collection("orders")
	orderService = services.NewOrderServiceImpl(orderCollection, ctx)
	OrderController = controllers.NewOrderController(orderService, ctx, bot, config.Admins)
	OrderRouteController = routes.NewOrderRouteController(OrderController)
	server = gin.Default()

}
func NewBot(conf config.Config) *tele.Bot {
	pref := tele.Settings{
		Token:  conf.Token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	b.Handle("/start", handlers.StartCommand)

	return b
}
