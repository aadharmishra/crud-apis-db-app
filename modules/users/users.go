package users

import (
	"crud-apis-db-app/modules/users/models"
	"crud-apis-db-app/shared"
	"fmt"
	"net/http"
	"strconv"

	"crud-apis-db-app/dal/users"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type UsersModule struct {
	Deps     *shared.Deps
	UsersDal users.UsersDal
}

func NewUsersModule(deps *shared.Deps) *UsersModule {
	return &UsersModule{
		Deps:     deps,
		UsersDal: users.NewUsersDal(deps),
	}
}

func (u *UsersModule) CreateNewUser(ctx *gin.Context) (int, *models.User, error) {

	var request *models.User
	err := ctx.ShouldBindBodyWith(&request, binding.JSON)
	if err != nil || request == nil {
		fmt.Println("bad request")
		return http.StatusBadRequest, nil, err
	}

	exists, err := u.UsersDal.IsUserAvailable(ctx, request.ID)

	if exists {
		fmt.Println("user already exists")
		return http.StatusOK, nil, err
	}

	err = u.UsersDal.InsertUser(ctx, request)
	if err != nil {
		fmt.Println("error while inserting record into DB")
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, request, err
}

func (u *UsersModule) FetchAllUsers(ctx *gin.Context) (int, *[]models.User, error) {
	result, err := u.UsersDal.SelectUsers(ctx)

	if result == nil || err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, &result, nil
}

func (u *UsersModule) FetchUser(ctx *gin.Context) (int, *models.User, error) {
	id := ctx.Param("id")

	intId, err := strconv.Atoi(id)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	result, err := u.UsersDal.SelectUsersById(ctx, intId)

	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, &result, nil
}

func (u *UsersModule) UpdateExistingUser(ctx *gin.Context) (int, *models.User, error) {
	var request *models.User
	err := ctx.ShouldBindBodyWith(&request, binding.JSON)
	if err != nil || request == nil {
		fmt.Println("bad request")
		return http.StatusBadRequest, nil, err
	}

	exists, err := u.UsersDal.IsUserAvailable(ctx, request.ID)

	if !exists {
		fmt.Println("user doesn't exist")
		return http.StatusOK, nil, err
	}

	updated, err := u.UsersDal.UpdateUserById(ctx, request)
	if err != nil {
		fmt.Println("error while inserting record into DB")
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, &updated, err
}

func (u *UsersModule) RemoveExistingUser(ctx *gin.Context) (int, bool, error) {
	id := ctx.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		return http.StatusBadRequest, false, err
	}

	deleted, err := u.UsersDal.DeleteUserById(ctx, intId)
	if err != nil || !deleted {
		return http.StatusInternalServerError, false, err
	}

	return http.StatusOK, true, nil
}
