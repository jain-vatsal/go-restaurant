package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jain-vatsal/go-restaurant/controllers"
	"github.com/jain-vatsal/go-restaurant/endpoints"
)

func FoodRoutes(foodRoutes *gin.Engine) {

	foodRoutes.GET(endpoints.Foods, controllers.GetFoods())
	foodRoutes.GET(endpoints.FoodId, controllers.GetFood())
	foodRoutes.POST(endpoints.Foods, controllers.CreateFood())
	foodRoutes.PATCH(endpoints.FoodId, controllers.UpdateFood())
}
