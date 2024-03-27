package models

import "time"

type Department struct {
	Id              int       `json:"id"`
	Name            string    `json:"name"`
	Location        string    `json:"location"`
	Email           string    `json:"email"`
	Budget          int       `json:"budget"`
	IsCentralBranch bool      `json:"isCentralBranch"`
	Description     string    `json:"description"`
	LastModified    time.Time `json:"lastModified,omitempty"`
}
