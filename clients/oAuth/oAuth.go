package oauth

import (
	"context"
	"crud-apis-db-app/config"
	"errors"
	"io"
	"net/http"
	"strings"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type IOAuthCfg interface {
	GetAuthCodeURL(ctx context.Context) (string, error)
	GetToken(ctx context.Context, code string) (string, error)
	UserInfo(ctx context.Context, url string, token string) (string, error)
	YoutubeSearch(ctx context.Context, url string, token string, queryParams string) (string, error)
	UploadToGoogleDrive(ctx context.Context, url string, token string, queryParams string, body string) (string, error)
}

type GoogleOAuthClientCfg struct {
	OauthCfg *oauth2.Config
}

var GoogleOAuth2Cfg = GoogleOAuthClientCfg{}

func NewOAuthClientCfg(config config.IConfig) (IOAuthCfg, error) {
	oAuthRawCfg := config.Get().OAuth.Web

	googleOAuthCfg := &oauth2.Config{
		ClientID:     oAuthRawCfg.Client_id,
		ClientSecret: oAuthRawCfg.Client_secret,
		RedirectURL:  oAuthRawCfg.Redirect_uris[0],
		Scopes:       oAuthRawCfg.Scopes,
		Endpoint:     google.Endpoint,
	}

	GoogleOAuth2Cfg = GoogleOAuthClientCfg{OauthCfg: googleOAuthCfg}

	return &GoogleOAuth2Cfg, nil
}

func (g *GoogleOAuthClientCfg) GetAuthCodeURL(ctx context.Context) (string, error) {
	var url string

	googleCfg := g.OauthCfg

	url = googleCfg.AuthCodeURL("randomstate")

	if url == "" {
		return "", errors.New("empty authcode url")
	}

	return url, nil
}

func (g *GoogleOAuthClientCfg) GetToken(ctx context.Context, code string) (string, error) {

	googleCfg := g.OauthCfg

	token, err := googleCfg.Exchange(ctx, code)
	if err != nil {
		return "", err
	}

	if token == nil {
		return "", errors.New("empty token received")
	}

	return token.AccessToken, nil
}

func (g *GoogleOAuthClientCfg) UserInfo(ctx context.Context, url string, token string) (string, error) {
	res, err := http.Get(url + token)
	if err != nil {
		return "", err
	}

	if res == nil {
		return "", errors.New("invalid response")
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	if data == nil {
		return "", errors.New("invalid response")
	}

	return string(data), nil
}

func (g *GoogleOAuthClientCfg) YoutubeSearch(ctx context.Context, url string, token string, queryParams string) (string, error) {
	finalUrl := url + token + queryParams
	res, err := http.Get(finalUrl)
	if err != nil {
		return "", err
	}

	if res == nil {
		return "", errors.New("invalid response")
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	if data == nil {
		return "", errors.New("invalid response")
	}

	return string(data), nil
}

func (g *GoogleOAuthClientCfg) UploadToGoogleDrive(ctx context.Context, url string, token string, queryParams string, body string) (string, error) {

	finalUrl := url + token + queryParams
	req, err := http.NewRequest("POST", finalUrl, strings.NewReader(body))
	if err != nil {
		return "", err
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(resBody), nil
}
