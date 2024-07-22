package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jain-vatsal/go-restaurant/controllers"
	"github.com/jain-vatsal/go-restaurant/endpoints"
)

func OrderRoutes(orderRoutes *gin.Engine) {

	orderRoutes.GET(endpoints.Orders, controllers.GetOrders())
	orderRoutes.GET(endpoints.OrderId, controllers.GetOrder())
	orderRoutes.POST(endpoints.Orders, controllers.CreateOrder())
	orderRoutes.PATCH(endpoints.OrderId, controllers.UpdateOrder())

}
