package controllers

import (
	"WhatsappOrderServer/models"
	"WhatsappOrderServer/services"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type OrderController struct {
	orderService services.OrderService
	ctx          context.Context
}

func NewOrderController(orderService services.OrderService, ctx context.Context) OrderController {
	return OrderController{orderService: orderService, ctx: ctx}
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
	message := "Order ID - " + orderId
	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "message": message})

}
