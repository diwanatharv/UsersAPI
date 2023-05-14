package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"awesomeProject3/api/service"
	"awesomeProject3/pkg/enums"
	"awesomeProject3/pkg/models"
)

func GetallUsers(e echo.Context) error {
	var resbody []models.User
	res, err := service.Getalluser(resbody)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, enums.Servererror)
	}
	return e.JSON(http.StatusOK, res)
}
