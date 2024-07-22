package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jain-vatsal/go-restaurant/controllers"
	"github.com/jain-vatsal/go-restaurant/endpoints"
)

func OrderItemRoutes(orderItemRoutes *gin.Engine) {

	orderItemRoutes.GET(endpoints.OrderItems, controllers.GetOrderItems())
	orderItemRoutes.GET(endpoints.OrderItemId, controllers.GetOrderItem())
	orderItemRoutes.POST(endpoints.OrderItems, controllers.CreateOrderItem())
	orderItemRoutes.PATCH(endpoints.OrderItemId, controllers.UpdateOrderItem())
	orderItemRoutes.GET(endpoints.AllItemsForOrderID, controllers.GetOrderItemsByOrderID())
}
