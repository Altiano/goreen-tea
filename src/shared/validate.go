package shared

import (
	"net/http"

	"github.com/go-playground/validator"
)

func Validate(req interface{}) error {
	v := validator.New()
	if err := v.Struct(req); err != nil {
		return MakeHTTPError(http.StatusBadRequest, err)
	}

	return nil
}
