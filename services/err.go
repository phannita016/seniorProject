package services

import (
	"github.com/phannita016/seniorProject/x/errs"
	"github.com/phannita016/seniorProject/x/libs"
)

type Error = errs.Error

func ErrService(err error) error {
	return &Error{Opt: libs.SERVICE, Err: err}
}
