package routes

import (
	"github.com/gin-gonic/gin"
	"log"
)

func SetUpRouter(port string) {

	router := gin.New()
	router.Use(gin.Logger())

	UserRoutes(router)
	MenuRoutes(router)
	TableRoutes(router)
	OrderItemRoutes(router)
	OrderRoutes(router)
	InvoiceRoutes(router)
	FoodRoutes(router)
	routerRunErr := router.Run(":" + port)
	if routerRunErr != nil {
		log.Fatal("Error while setting up router: ", routerRunErr.Error())
	}
}
