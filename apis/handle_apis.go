package apis

import (
	"time"

	"github.com/gofiber/fiber/v2"
	jwx "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
	"github.com/phannita016/seniorProject/x/libs"
)

const APIVersion1 = "/api/v1"

var Version = "0.0.1"

var Fc = fiber.Config{
	ServerHeader: Version,
	BodyLimit:    10 * 1024 * 1024, // 10 MB
	ReadTimeout:  30 * time.Second,
	WriteTimeout: 30 * time.Second,
	IdleTimeout:  10 * time.Second,
	// ErrorHandler: FiberErrorHandler,
}

func Authorization() fiber.Handler {
	return jwx.New(
		jwx.Config{
			Filter: func(c *fiber.Ctx) bool {
				switch c.Path() {
				case APIVersion1 + "/login", APIVersion1 + "/health", "/health":
					return true
				default:
					return false
				}
			},
			SigningKey:  []byte{},
			ContextKey:  libs.ContextKEY,
			TokenLookup: "header:Authorization",
			AuthScheme:  "Bearer",
			SuccessHandler: func(c *fiber.Ctx) error {
				if user, ok := c.Locals(libs.ContextKEY).(*jwt.Token); ok {
					if claims, ok := user.Claims.(jwt.MapClaims); ok {
						c.Locals(libs.ContextUSERID, claims["jti"])
						c.Locals(libs.ContextUSEREMAIL, claims["sub"])
					}
				}
				return c.Next()
			},
			ErrorHandler: func(c *fiber.Ctx, err error) error {
				return errUnauthorized(err)
			},
		},
	)
}

func SetHandle(name string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Locals(libs.Handle, name)
		return c.Next()
	}
}
