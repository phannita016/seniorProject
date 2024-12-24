package apps

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/phannita016/seniorProject/apis"
	"github.com/phannita016/seniorProject/services"
)

func NewAPIs(serviveUser services.User) *fiber.App {
	f := fiber.New(apis.Fc)
	// f.Use(limiter.New(apis.))
	f.Use(recover.New())
	f.Use(cors.New())
	f.Use(compress.New())
	f.Use(requestid.New())

	f.Use(apis.Authorization())

	g := f.Group(apis.APIVersion1)
	apis.NewHandleLogin(g, serviveUser)

	return f
}
