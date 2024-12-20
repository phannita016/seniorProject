package libs

import (
	"context"
	"time"

	"github.com/google/uuid"
)

var LoginSecret = uuid.NewString()

const (
	ContextUSERID    = "USER_ID"
	ContextRequestID = "requestid"
	ContextUSEREMAIL = "USER_EMAIL"
)

const (
	ContextKEY = "ctk"
)

const (
	Handle = "HANDLE"

	Login = "Login"
)

const (
	SERVICE      = "SERVICE"
	UNAUTHORIZED = "UNAUTHORIZED"
	VALIDATION   = "VALIDATION"
)

type Context func() (context.Context, context.CancelFunc)

func ContextTimeout(seconds ...int) Context {
	second := 5
	if len(seconds) > 0 {
		second = seconds[len(seconds)-1]
	}

	return func() (context.Context, context.CancelFunc) {
		return context.WithTimeout(context.Background(), time.Second*time.Duration(second))
	}
}
