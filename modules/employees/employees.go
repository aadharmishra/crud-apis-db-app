package employees

import (
	"crud-apis-db-app/dal/employees"
	"crud-apis-db-app/modules/employees/models"
	"crud-apis-db-app/shared"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type EmployeesModule struct {
	Deps         *shared.Deps
	EmployeesDal employees.EmployeesDal
}

func NewEmployeesModule(deps *shared.Deps) *EmployeesModule {
	return &EmployeesModule{
		Deps:         deps,
		EmployeesDal: employees.NewEmployeesDal(deps),
	}
}

func (e *EmployeesModule) CreateNewEmployee(ctx *gin.Context) (int, *[]models.Employee, error) {
	var request *[]models.Employee
	err := ctx.ShouldBindBodyWith(&request, binding.JSON)
	if err != nil || request == nil {
		fmt.Println("bad request")
		return http.StatusBadRequest, nil, err
	}

	err = e.EmployeesDal.InsertEmployee(ctx, request)
	if err != nil {
		fmt.Println("error while inserting record into DB")
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, request, err
}

func (e *EmployeesModule) FetchAllEmployees(ctx *gin.Context) (int, *[]models.Employee, error) {
	result, err := e.EmployeesDal.SelectEmployees(ctx)

	if result == nil || err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, result, nil
}

func (e *EmployeesModule) FetchEmployee(ctx *gin.Context) (int, *models.Employee, error) {
	id := ctx.Param("id")

	intId, err := strconv.Atoi(id)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	result, err := e.EmployeesDal.SelectEmployeeById(ctx, intId)

	if result == nil || err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, result, nil
}

func (e *EmployeesModule) UpdateExistingEmployees(ctx *gin.Context) (int, *[]models.Employee, error) {
	var request *[]models.Employee
	err := ctx.ShouldBindBodyWith(&request, binding.JSON)
	if err != nil || request == nil {
		fmt.Println("bad request")
		return http.StatusBadRequest, nil, err
	}

	result, err := e.EmployeesDal.UpdateEmployeesById(ctx, request)
	if err != nil {
		fmt.Println("error while updating records in DB")
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, result, err
}

func (u *EmployeesModule) RemoveExistingEmployees(ctx *gin.Context) (int, bool, error) {
	strIds := ctx.Query("ids")
	if len(strIds) == 0 {
		return http.StatusBadRequest, false, errors.New("bad request")
	}

	ids := strings.Split(strIds, ",")

	var idList []int
	for _, id := range ids {
		intId, err := strconv.Atoi(id)
		if err != nil {
			return http.StatusBadRequest, false, err
		}
		idList = append(idList, intId)
	}

	deleted, err := u.EmployeesDal.DeleteEmployeeById(ctx, idList)
	if err != nil || !deleted {
		return http.StatusInternalServerError, false, err
	}

	return http.StatusOK, true, nil
}
