package departments

import (
	"crud-apis-db-app/dal/departments"
	"crud-apis-db-app/modules/departments/models"
	"crud-apis-db-app/shared"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type DepartmentsModule struct {
	Deps           *shared.Deps
	DepartmentsDal departments.DepartmentsDal
}

func NewDepartmentsModule(deps *shared.Deps) *DepartmentsModule {
	return &DepartmentsModule{
		Deps:           deps,
		DepartmentsDal: departments.NewDepartmentsDal(deps),
	}
}

func (e *DepartmentsModule) CreateNewDepartment(ctx *gin.Context) (int, *[]models.Department, error) {
	var request *[]models.Department
	err := ctx.ShouldBindBodyWith(&request, binding.JSON)
	if err != nil || request == nil {
		fmt.Println("bad request")
		return http.StatusBadRequest, nil, err
	}

	err = e.DepartmentsDal.InsertDepartment(ctx, request)
	if err != nil {
		fmt.Println("error while inserting record into DB")
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, request, err
}

func (e *DepartmentsModule) FetchAllDepartments(ctx *gin.Context) (int, *[]models.Department, error) {
	result, err := e.DepartmentsDal.SelectDepartments(ctx)

	if result == nil || err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, result, nil
}

func (e *DepartmentsModule) FetchDepartment(ctx *gin.Context) (int, *models.Department, error) {
	id := ctx.Param("id")

	intId, err := strconv.Atoi(id)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	result, err := e.DepartmentsDal.SelectDepartmentById(ctx, intId)

	if result == nil || err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, result, nil
}

func (e *DepartmentsModule) UpdateExistingDepartments(ctx *gin.Context) (int, *[]models.Department, error) {
	var request *[]models.Department
	err := ctx.ShouldBindBodyWith(&request, binding.JSON)
	if err != nil || request == nil {
		fmt.Println("bad request")
		return http.StatusBadRequest, nil, err
	}

	result, err := e.DepartmentsDal.UpdateDepartmentsById(ctx, request)
	if err != nil {
		fmt.Println("error while updating records in DB")
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, result, err
}

func (u *DepartmentsModule) RemoveExistingDepartments(ctx *gin.Context) (int, bool, error) {
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

	deleted, err := u.DepartmentsDal.DeleteDepartmentById(ctx, idList)
	if err != nil || !deleted {
		return http.StatusInternalServerError, false, err
	}

	return http.StatusOK, true, nil
}
