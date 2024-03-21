package models

import "time"

type User struct {
	ID          int            `json:"id"`
	Username    string         `json:"username"`
	Email       string         `json:"email"`
	Age         int            `json:"age"`
	IsAdmin     bool           `json:"isadmin"`
	LastLogin   *time.Time     `json:"lastlogin,omitempty"`
	Preferences map[string]any `json:"preferences,omitempty"`
}
