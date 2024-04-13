package oauth

import (
	"crud-apis-db-app/shared"
	"net/http"

	oauth "crud-apis-db-app/modules/oAuth"

	"github.com/gin-gonic/gin"
)

type OAuthService struct {
	oauth *oauth.OAuthModule
}

func NewOAuthService(deps *shared.Deps) *OAuthService {
	return &OAuthService{
		oauth: oauth.NewOAuthModule(deps),
	}
}

func (o *OAuthService) Login(ctx *gin.Context) {
	var err error
	statusCode, url, err := o.oauth.LoginUser(ctx)
	if err != nil || url == "" || statusCode != http.StatusOK {
		ctx.JSON(statusCode, map[string]interface{}{"error": "internal error"})
		return
	}
	ctx.Redirect(http.StatusSeeOther, url)
}

func (o *OAuthService) Callback(ctx *gin.Context) {
	var err error
	statusCode, res, err := o.oauth.TriggerCallback(ctx)
	if err != nil || res == nil || statusCode != http.StatusOK {
		ctx.JSON(statusCode, map[string]interface{}{"error": "internal error"})
		return
	}
	ctx.JSON(statusCode, res)
}

func (o *OAuthService) YoutubeSearch(ctx *gin.Context) {
	var err error
	statusCode, res, err := o.oauth.YoutubeSearchData(ctx)
	if err != nil || res == nil || statusCode != http.StatusOK {
		ctx.JSON(statusCode, map[string]interface{}{"error": "internal error"})
		return
	}
	ctx.JSON(statusCode, res)
}

func (o *OAuthService) CreateDriveFile(ctx *gin.Context) {
	var err error
	statusCode, res, err := o.oauth.CreateNewMultipartDriveFile(ctx)
	if err != nil || res == nil || statusCode != http.StatusOK {
		ctx.JSON(statusCode, map[string]interface{}{"error": "internal error"})
		return
	}
	ctx.JSON(statusCode, res)
}
