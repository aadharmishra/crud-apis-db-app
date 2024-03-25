package employees

import (
	"crud-apis-db-app/shared"
	"net/http"

	"crud-apis-db-app/modules/employees"

	"github.com/gin-gonic/gin"
)

type EmployeesService struct {
	employees *employees.EmployeesModule
}

func NewEmployeesService(deps *shared.Deps) *EmployeesService {
	return &EmployeesService{
		employees: employees.NewEmployeesModule(deps),
	}
}

func (e *EmployeesService) CreateEmployee(ctx *gin.Context) {
	var err error
	statusCode, res, err := e.employees.CreateNewEmployee(ctx)
	if err != nil || res == nil || statusCode != http.StatusOK {
		ctx.JSON(statusCode, map[string]interface{}{"error": "internal error"})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (e *EmployeesService) GetAllEmployees(ctx *gin.Context) {
	var err error
	statusCode, res, err := e.employees.FetchAllEmployees(ctx)
	if err != nil || res == nil || statusCode != http.StatusOK {
		ctx.JSON(statusCode, map[string]interface{}{"error": "internal error"})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (e *EmployeesService) GetEmployee(ctx *gin.Context) {
	var err error
	statusCode, res, err := e.employees.FetchEmployee(ctx)
	if err != nil || res == nil || statusCode != http.StatusOK {
		ctx.JSON(statusCode, map[string]interface{}{"error": "internal error"})
		return
	}
	ctx.JSON(http.StatusOK, res)

}

func (e *EmployeesService) UpdateEmployees(ctx *gin.Context) {
	var err error
	statusCode, res, err := e.employees.UpdateExistingEmployees(ctx)
	if err != nil || res == nil || statusCode != http.StatusOK {
		ctx.JSON(statusCode, map[string]interface{}{"error": "internal error"})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (e *EmployeesService) RemoveEmployees(ctx *gin.Context) {
	var err error
	statusCode, removed, err := e.employees.RemoveExistingEmployees(ctx)
	if err != nil || !removed || statusCode != http.StatusOK {
		ctx.JSON(statusCode, map[string]interface{}{"error": "internal error"})
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{"removed": "success"})
}
