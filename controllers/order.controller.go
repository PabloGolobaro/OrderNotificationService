package controllers

import (
	"WhatsappOrderServer/models"
	"WhatsappOrderServer/services"
	"WhatsappOrderServer/utils"
	"context"
	"github.com/gin-gonic/gin"
	tele "gopkg.in/telebot.v3"
	"net/http"
)

type OrderController struct {
	orderService services.OrderService
	ctx          context.Context
	bot          *tele.Bot
	admins       []string
}

func NewOrderController(orderService services.OrderService, ctx context.Context, bot *tele.Bot, admins []string) OrderController {
	return OrderController{orderService: orderService, ctx: ctx, bot: bot, admins: admins}
}
func (oc *OrderController) SaveOrder(ctx *gin.Context) {

	var order *models.Order

	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	orderId, err := oc.orderService.SaveOrder(order)
	if err != nil {
		errMessage := "Извините, что-то пошло не так, для заказа позвоните по номеру +7 (965) 989-89-88, ждем вас"
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": errMessage})
		return
	}
	order.OrderId = orderId
	err = utils.SendToAdmins(oc.bot, oc.admins, order)
	if err != nil {
		errMessage := "Извините, что-то пошло не так, для заказа позвоните по номеру +7 (965) 989-89-88, ждем вас"
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": errMessage})
		return
	}
	message := "Order ID - " + orderId
	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "message": message})

}
