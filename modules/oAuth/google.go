package oauth

import (
	"bytes"
	"crud-apis-db-app/modules/oAuth/models"
	"crud-apis-db-app/shared"
	"encoding/json"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"time"

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

	err = o.Deps.Database.RedisDb.Create(ctx, "token", token, 24*time.Hour)
	if err != nil {
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

func (o *OAuthModule) YoutubeSearchData(ctx *gin.Context) (int, map[string]interface{}, error) {
	token, err := o.Deps.Database.RedisDb.Read(ctx, "token")
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	if token == "" {
		return http.StatusInternalServerError, nil, errors.New("internal server error")
	}

	url := o.Deps.Config.Get().Google.YoutubeSearchUrl

	strQueryParams := "&part=snippet&forMine=true&type=video"

	res, err := o.Deps.ClientConfigs.OAuthClientCfg.YoutubeSearch(ctx, url, token, strQueryParams)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	if res == "" {
		return http.StatusInternalServerError, nil, errors.New("internal server error")
	}

	var jsonData map[string]interface{}
	err = json.Unmarshal([]byte(res), &jsonData)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, jsonData, nil
}

func (o *OAuthModule) CreateNewDriveFile(ctx *gin.Context) (int, map[string]interface{}, error) {

	file, _, err := ctx.Request.FormFile("file")
	if err != nil {
		return http.StatusBadRequest, nil, err
	}
	defer file.Close()

	fileData, err := io.ReadAll(file)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	token, err := o.Deps.Database.RedisDb.Read(ctx, "token")
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	if token == "" {
		return http.StatusInternalServerError, nil, errors.New("internal server error")
	}

	url := o.Deps.Config.Get().Google.GoogleDriveUploadUrl

	strQueryParams := "&uploadType=media"

	res, err := o.Deps.ClientConfigs.OAuthClientCfg.UploadToGoogleDrive(ctx, url, token, strQueryParams, string(fileData))
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	if res == "" {
		return http.StatusInternalServerError, nil, errors.New("internal server error")
	}

	var jsonData map[string]interface{}
	err = json.Unmarshal([]byte(res), &jsonData)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, jsonData, nil
}

func (o *OAuthModule) CreateNewMultipartDriveFile(ctx *gin.Context) (int, map[string]interface{}, error) {

	// Access uploaded file
	file, err := ctx.FormFile("file")
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Media part for the uploaded file
	mediaPart, err := writer.CreatePart(textproto.MIMEHeader{
		"Content-Type": {file.Header.Get("Content-Type")},
	})
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	openedFile, err := file.Open()
	if err != nil {
		return http.StatusBadRequest, nil, err
	}
	defer openedFile.Close()

	_, err = io.Copy(mediaPart, openedFile)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	writer.Close()

	// Access token retrieval (unchanged)
	token, err := o.Deps.Database.RedisDb.Read(ctx, "token")
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	if token == "" {
		return http.StatusInternalServerError, nil, errors.New("internal server error")
	}

	// Construct HTTP request for Google Drive upload
	req, err := http.NewRequest("POST", "https://www.googleapis.com/upload/drive/v3/files?access_token="+token+"&uploadType=multipart", body)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.ContentLength = int64(body.Len())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}
	defer resp.Body.Close()

	res, err := io.ReadAll(resp.Body)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	var jsonData map[string]interface{}
	err = json.Unmarshal(res, &jsonData)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	// Potentially use `request` variable for Google Drive API call based on its requirements

	return http.StatusOK, jsonData, nil
}
