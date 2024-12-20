package apis

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/phannita016/seniorProject/x/errs"
	"github.com/phannita016/seniorProject/x/libs"
)

const (
	UNAUTHORIZED = libs.UNAUTHORIZED
	VALIDATION   = libs.VALIDATION

	KeyBodyParser = "BodyParser"
)

var (
	ErrBodyParser = errBodyParser()
)

type APIError = errs.APIError
type ErrorValidation = errs.ErrorValidation
type ErrorValidator = errs.ErrorValidator

func errUnauthorized(err error) error {
	return &APIError{Code: fiber.StatusUnauthorized, Opt: UNAUTHORIZED, Err: err}
}

func errBodyParser() error {
	return &APIError{Code: fiber.StatusNotAcceptable, Opt: VALIDATION, Err: errors.New(KeyBodyParser)}
}

func ErrValidator(errs []*ErrorValidation) error {
	return &ErrorValidator{Code: fiber.StatusBadRequest, Opt: VALIDATION, ErrorValidation: errs}
}
