package helpers

import (
	"github.com/fydhfzh/assignment-3/pkg/errs"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateStruct(payload interface{}) errs.ErrMessage {
	err := validate.Struct(payload)
	if err != nil {
		return errs.NewBadRequest("bad request body")
	}

	return nil
}