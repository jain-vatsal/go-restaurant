package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jain-vatsal/go-restaurant/controllers"
	"github.com/jain-vatsal/go-restaurant/endpoints"
)

func TableRoutes(tableRoutes *gin.Engine) {

	tableRoutes.GET(endpoints.Tables, controllers.GetTables())
	tableRoutes.GET(endpoints.TableId, controllers.GetTable())
	tableRoutes.POST(endpoints.Tables, controllers.CreateTable())
	tableRoutes.PATCH(endpoints.TableId, controllers.UpdateTable())

}
