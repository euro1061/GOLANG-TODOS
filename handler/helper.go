package handler

import (
	"GOLANG-TODOS/errs"
	"GOLANG-TODOS/service"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func handleError(c *fiber.Ctx, err error) error {
	switch e := err.(type) {
	case errs.AppError:
		return c.Status(e.Code).JSON(e.JSONMessage)
	case error:
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": e.Error(),
		})
	}
	return nil
}

var validate = validator.New()

func validateStruct(body any) []*service.IError {
	var errors []*service.IError

	err := validate.Struct(body)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element service.IError
			element.Field = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}

	return errors
}
