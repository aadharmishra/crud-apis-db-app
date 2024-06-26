package models

type UserInfo struct {
	Id             string `json:"id"`
	Email          string `json:"email"`
	Verified_email bool   `json:"verified_email"`
	Name           string `json:"name"`
	Given_name     string `json:"given_name"`
	Family_name    string `json:"family_name"`
	Picture        string `json:"picture"`
	Locale         string `json:"locale"`
}

type UploadRequest struct {
	Name        string `json:"name"`
	MimeType    string `json:"mimeType"`
	Description string `json:"description"`
}
