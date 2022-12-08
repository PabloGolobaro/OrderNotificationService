package routes

import (
	"WhatsappOrderServer/controllers"
	"github.com/gin-gonic/gin"
)

type OrderRouteController struct {
	orderController controllers.OrderController
}

func NewOrderRouteController(orderController controllers.OrderController) OrderRouteController {
	return OrderRouteController{orderController: orderController}
}
func (rc *OrderRouteController) AuthRoute(rg *gin.RouterGroup) {
	rg.POST("order", rc.orderController.SaveOrder)
}
