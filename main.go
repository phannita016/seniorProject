package main

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/phannita016/seniorProject/config"
)

func main() {
	app := fiber.New()

	db := config.ConnectDB()
	defer func() {
		if err := db.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	if err := app.Listen(":6000"); err != nil {
		panic(err)
	}
}
