package services

import (
	"time"

	"github.com/phannita016/seniorProject/dtos"
	"github.com/phannita016/seniorProject/stores"
	"github.com/phannita016/seniorProject/x/libs"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type User interface {
	Service[dtos.User]
	Mocking(username, password string) error
}

type UserStore stores.Store[dtos.User]

type user struct {
	Service[dtos.User]
	store   UserStore
	context libs.Context
}

func NewUser(store UserStore, ctx libs.Context) User {
	return &user{
		Service: nil,
		store:   store,
		context: ctx,
	}
}

func (u *user) Create(dto dtos.UserDtos) error {
	ctx, cancel := u.context()
	defer cancel()

	bt, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)
	if err != nil {
		return ErrService(err)
	}

	var prod = dtos.User{
		ID:        primitive.NewObjectID(),
		Email:     dto.Email,
		Password:  string(bt),
		CreatedAt: time.Now(),
	}

	return u.store.Create(ctx, prod)
}

func (u *user) Mocking(username, password string) error {
	ctx, cancel := u.context()
	defer cancel()

	finds, err := u.store.FindAll(ctx)
	if err != nil {
		return err
	}
	if len(finds) > 0 {
		return nil
	}

	var create = func() error {
		return u.Create(dtos.UserDtos{
			Email:    username,
			Password: password,
		})
	}

	find, err := u.store.FindOneKeyValue(ctx, "email", username)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			return err
		}
		return create()
	}

	if find == nil || find.ID.IsZero() {
		return create()
	}

	return nil
}
