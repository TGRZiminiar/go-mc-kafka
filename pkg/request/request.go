package request

import (
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type (
	contextWrapperService interface {
		Bind(data any) error
	}

	contextWrapepr struct {
		Context  echo.Context
		validaor *validator.Validate
	}
)

func NewContextWrapper(ctx echo.Context) contextWrapperService {
	return &contextWrapepr{
		Context:  ctx,
		validaor: validator.New(),
	}
}

func (c contextWrapepr) Bind(data any) error {
	if err := c.Context.Bind(data); err != nil {
		log.Printf("Error: Bind data failed: %s", err.Error())
	}

	if err := c.validaor.Struct(data); err != nil {
		log.Printf("Error: Validate data failed: %s", err.Error())
	}
	return nil
}
