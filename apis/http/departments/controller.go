package departments

import (
	"crud-apis-db-app/shared"
	"net/http"

	"crud-apis-db-app/modules/departments"

	"github.com/gin-gonic/gin"
)

type DepartmentsService struct {
	departments *departments.DepartmentsModule
}

func NewDepartmentsService(deps *shared.Deps) *DepartmentsService {
	return &DepartmentsService{
		departments: departments.NewDepartmentsModule(deps),
	}
}

func (d *DepartmentsService) CreateDepartment(ctx *gin.Context) {
	var err error
	statusCode, res, err := d.departments.CreateNewDepartment(ctx)
	if err != nil || res == nil || statusCode != http.StatusOK {
		ctx.JSON(statusCode, map[string]interface{}{"error": "internal error"})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (d *DepartmentsService) GetAllDepartments(ctx *gin.Context) {
	var err error
	statusCode, res, err := d.departments.FetchAllDepartments(ctx)
	if err != nil || res == nil || statusCode != http.StatusOK {
		ctx.JSON(statusCode, map[string]interface{}{"error": "internal error"})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (d *DepartmentsService) GetDepartment(ctx *gin.Context) {
	var err error
	statusCode, res, err := d.departments.FetchDepartment(ctx)
	if err != nil || res == nil || statusCode != http.StatusOK {
		ctx.JSON(statusCode, map[string]interface{}{"error": "internal error"})
		return
	}
	ctx.JSON(http.StatusOK, res)

}

func (d *DepartmentsService) UpdateDepartments(ctx *gin.Context) {
	var err error
	statusCode, res, err := d.departments.UpdateExistingDepartments(ctx)
	if err != nil || res == nil || statusCode != http.StatusOK {
		ctx.JSON(statusCode, map[string]interface{}{"error": "internal error"})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (d *DepartmentsService) RemoveDepartments(ctx *gin.Context) {
	var err error
	statusCode, removed, err := d.departments.RemoveExistingDepartments(ctx)
	if err != nil || !removed || statusCode != http.StatusOK {
		ctx.JSON(statusCode, map[string]interface{}{"error": "internal error"})
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{"removed": "success"})
}
