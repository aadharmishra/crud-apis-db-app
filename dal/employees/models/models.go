package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Employee struct {
	ID           int                `bson:"id"`
	Username     string             `bson:"username"`
	Email        string             `bson:"email"`
	Age          int                `bson:"age"`
	IsAdmin      bool               `bson:"isAdmin"`
	Dob          string             `bson:"dob"`
	Details      string             `bson:"details"`
	LastModified primitive.DateTime `bson:"lastModified,omitempty"`
}
