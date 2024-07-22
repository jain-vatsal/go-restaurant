package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jain-vatsal/go-restaurant/controllers"
	"github.com/jain-vatsal/go-restaurant/endpoints"
)

func UserRoutes(userRoutes *gin.Engine) {

	userRoutes.GET(endpoints.Users, controllers.GetUsers())
	userRoutes.GET(endpoints.UserId, controllers.GetUser())
	userRoutes.POST(endpoints.UserSignUp, controllers.UserSignUp())
	userRoutes.POST(endpoints.UserLogin, controllers.UserLogin())

}
