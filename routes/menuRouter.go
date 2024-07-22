package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jain-vatsal/go-restaurant/controllers"
	"github.com/jain-vatsal/go-restaurant/endpoints"
)

func MenuRoutes(menuRoutes *gin.Engine) {

	menuRoutes.GET(endpoints.Menus, controllers.GetMenus())
	menuRoutes.GET(endpoints.MenuId, controllers.GetMenu())
	menuRoutes.POST(endpoints.Menus, controllers.CreateMenu())
	menuRoutes.PATCH(endpoints.MenuId, controllers.UpdateMenu())

}
