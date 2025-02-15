package utils

import "github.com/labstack/echo/v4"

func HandleGormError(c echo.Context, err error) error {
	return c.JSON(400, echo.Map{"message": err.Error()})
}
