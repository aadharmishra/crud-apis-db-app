package users

import (
	"crud-apis-db-app/modules/users"
	"crud-apis-db-app/shared"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UsersService struct {
	users *users.UsersModule
}

func NewUsersService(deps *shared.Deps) *UsersService {
	return &UsersService{
		users: users.NewUsersModule(deps),
	}
}

func (u *UsersService) CreateUser(ctx *gin.Context) {
	var err error
	statusCode, res, err := u.users.CreateNewUser(ctx)
	if err != nil || res == nil || statusCode != http.StatusOK {
		ctx.JSON(statusCode, map[string]interface{}{"error": "internal error"})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (u *UsersService) GetAllUsers(ctx *gin.Context) {
	var err error
	statusCode, res, err := u.users.FetchAllUsers(ctx)
	if err != nil || res == nil || statusCode != http.StatusOK {
		ctx.JSON(statusCode, map[string]interface{}{"error": "internal error"})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (u *UsersService) GetUser(ctx *gin.Context) {
	var err error
	statusCode, res, err := u.users.FetchUser(ctx)
	if err != nil || res == nil || statusCode != http.StatusOK {
		ctx.JSON(statusCode, map[string]interface{}{"error": "internal error"})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (u *UsersService) UpdateUser(ctx *gin.Context) {
	var err error
	statusCode, res, err := u.users.UpdateExistingUser(ctx)
	if err != nil || res == nil || statusCode != http.StatusOK {
		ctx.JSON(statusCode, map[string]interface{}{"error": "internal error"})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (u *UsersService) RemoveUser(ctx *gin.Context) {
	var err error
	statusCode, removed, err := u.users.RemoveExistingUser(ctx)
	if err != nil || !removed || statusCode != http.StatusOK {
		ctx.JSON(statusCode, map[string]interface{}{"error": "internal error"})
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{"updated": "success"})
}
