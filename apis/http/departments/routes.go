package departments

import (
	"crud-apis-db-app/shared"

	"github.com/gin-gonic/gin"
)

func NewDepartmentsRouter(router *gin.Engine, deps *shared.Deps) {
	bindRoutes(router, deps)
}

func bindRoutes(router *gin.Engine, deps *shared.Deps) {
	service := NewDepartmentsService(deps)
	routerApi := router.Group("/departments")
	{
		routerApi.POST("/create", service.CreateDepartment)
		routerApi.GET("/get", service.GetAllDepartments)
		routerApi.GET("/get/:id", service.GetDepartment)
		routerApi.PUT("/update", service.UpdateDepartments)
		routerApi.DELETE("/remove", service.RemoveDepartments)
	}
}
