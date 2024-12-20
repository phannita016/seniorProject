package apis

import (
	"github.com/gofiber/fiber/v2"
	"github.com/phannita016/seniorProject/dtos"
	"github.com/phannita016/seniorProject/services"
	"github.com/phannita016/seniorProject/x/libs"
)

type login struct {
	service services.User
}

func NewHandleLogin(f fiber.Router, service services.User) {
	srv := login{
		service: service,
	}

	f.Use(SetHandle(libs.Login))
	f.Post("/login", HandleBodyParser(srv.Login))
}

func (l *login) Login(req dtos.UserDtos) (*dtos.UserToken, error) {
	return nil, nil
}
