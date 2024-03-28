package oauth

import (
	"crud-apis-db-app/modules/oAuth/models"
	"crud-apis-db-app/shared"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OAuthModule struct {
	Deps *shared.Deps
}

func NewOAuthModule(deps *shared.Deps) *OAuthModule {
	return &OAuthModule{
		Deps: deps,
	}
}

func (o *OAuthModule) LoginUser(ctx *gin.Context) (int, string, error) {

	url, err := o.Deps.ClientConfigs.OAuthClientCfg.GetAuthCodeURL(ctx)

	if err != nil {
		return http.StatusInternalServerError, "", err
	}

	return http.StatusOK, url, nil
}

func (o *OAuthModule) TriggerCallback(ctx *gin.Context) (int, *models.UserInfo, error) {
	state := ctx.QueryArray("state")[0]
	if state != "randomstate" {
		return http.StatusBadRequest, nil, errors.New("bad request")
	}

	code := ctx.QueryArray("code")[0]
	if code == "" {
		return http.StatusBadRequest, nil, errors.New("bad request")
	}

	token, err := o.Deps.ClientConfigs.OAuthClientCfg.GetToken(ctx, code)
	if err != nil {
		return http.StatusInternalServerError, nil, errors.New("internal server error")
	}

	if token == "" {
		return http.StatusInternalServerError, nil, errors.New("internal server error")
	}

	url := o.Deps.Config.Get().Google.UserInfoUrl

	res, err := o.Deps.ClientConfigs.OAuthClientCfg.UserInfo(ctx, url, token)
	if err != nil {
		return http.StatusInternalServerError, nil, errors.New("internal server error")
	}

	if res == "" {
		return http.StatusInternalServerError, nil, errors.New("internal server error")
	}

	var data models.UserInfo
	err = json.Unmarshal([]byte(res), &data)
	if err != nil {
		return http.StatusInternalServerError, nil, errors.New("internal server error")
	}

	return http.StatusOK, &data, nil
}
