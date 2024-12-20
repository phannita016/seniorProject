package apis

import (
	"log/slog"
	"reflect"

	"github.com/gofiber/fiber/v2"
)

func Parser[T any](t *T, parser func(any) error) error {
	if err := parser(t); err != nil {
		slog.Error("Parser", slog.Any("func", reflect.TypeOf(parser).Name()), slog.Any("err", err))
		return ErrBodyParser
	}

	if err := Validation(t); len(err) > 0 {
		return ErrValidator(err)
	}

	return nil
}

func HandleBodyParser[REQUEST, RESPONSE any](handle func(REQUEST) (RESPONSE, error)) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var request REQUEST
		if err := Parser(&request, c.BodyParser); err != nil {
			return err
		}

		res, err := handle(request)
		if err != nil {
			return err
		}

		return c.JSON(res)
	}
}
