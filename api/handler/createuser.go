package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"

	"awesomeProject3/api/service"
	"awesomeProject3/pkg/enums"
	"awesomeProject3/pkg/models"
)

func CreateUser(e echo.Context) error {
	var reqbody models.User
	err := json.NewDecoder(e.Request().Body).Decode(&reqbody)
	if err != nil {
		return e.JSON(http.StatusBadRequest, enums.Faileddecode)
	}
	v := validator.New()
	err = v.Struct(&reqbody)
	if err != nil {
		return e.JSON(http.StatusBadRequest, enums.Validationerror)
	}
	res, err := service.Createuser(reqbody)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, enums.Servererror)
	}
	return e.JSON(http.StatusOK, res)
}
