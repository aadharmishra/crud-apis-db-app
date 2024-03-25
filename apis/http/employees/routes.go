package employees

import (
	"crud-apis-db-app/shared"

	"github.com/gin-gonic/gin"
)

func NewEmployeesRouter(router *gin.Engine, deps *shared.Deps) {
	bindRoutes(router, deps)
}

func bindRoutes(router *gin.Engine, deps *shared.Deps) {
	service := NewEmployeesService(deps)
	routerApi := router.Group("/employees")
	{
		routerApi.POST("/create", service.CreateEmployee)
		routerApi.GET("/get", service.GetAllEmployees)
		routerApi.GET("/get/:id", service.GetEmployee)
		routerApi.PUT("/update", service.UpdateEmployees)
		routerApi.DELETE("/remove", service.RemoveEmployees)
	}
}
