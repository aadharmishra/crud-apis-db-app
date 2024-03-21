package users

import (
	"crud-apis-db-app/shared"

	"github.com/gin-gonic/gin"
)

func NewUsersRouter(router *gin.Engine, deps *shared.Deps) {
	bindRoutes(router, deps)
}

func bindRoutes(router *gin.Engine, deps *shared.Deps) {
	service := NewUsersService(deps)
	routerApi := router.Group("/users/")
	{
		routerApi.POST("/create", service.CreateUser)
		routerApi.GET("/get", service.GetAllUsers)
		routerApi.GET("/get/:id", service.GetUser)
		routerApi.PUT("/update", service.UpdateUser)
		routerApi.DELETE("remove/:id", service.RemoveUser)
	}
}
