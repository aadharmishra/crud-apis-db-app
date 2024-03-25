package models

import "time"

type Employee struct {
	ID           int       `json:"id"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	Age          int       `json:"age"`
	IsAdmin      bool      `json:"isAdmin"`
	Dob          string    `json:"dob"`
	Details      string    `json:"details"`
	LastModified time.Time `json:"lastModified,omitempty"`
}
