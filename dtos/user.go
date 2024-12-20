package dtos

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `json:"-" bson:"_id"`
	Email     string             `json:"email" bson:"email"`
	Password  string             `json:"-" bson:"password"`
	CreatedAt time.Time          `json:"-" bson:"create_at"`
	UpdateAt  time.Time          `json:"-" bson:"update_at"`
	DeleteAt  time.Time          `json:"-" bson:"delete_at"`
}

type UserDtos struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserToken struct {
	UserID  primitive.ObjectID `json:"-"`
	Message string             `json:"message"`
	Token   string             `json:"token"`
}
