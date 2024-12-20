package errs

import "encoding/json"

type Errs interface {
	Error() string
}

type Error struct {
	Opt string
	Err error
}

func (e *Error) Error() string {
	return e.Err.Error()
}

type APIError struct {
	Code int
	Opt  string
	Err  error
}

func (e *APIError) Error() string {
	return e.Err.Error()
}

type (
	ErrorValidator struct {
		Code            int
		Opt             string
		ErrorValidation []*ErrorValidation
	}

	ErrorValidation struct {
		Tag     string `json:"tag"`
		Field   string `json:"field"`
		Message string `json:"message"`
	}
)

func (e *ErrorValidator) Error() string {
	bt, err := json.Marshal(e.ErrorValidation)
	if err != nil {
		return err.Error()
	}

	return string(bt)
}
