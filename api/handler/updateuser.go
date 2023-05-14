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

func UpdateUser(e echo.Context) error {
	key := e.QueryParam("id")

	var reqBody models.User
	err1 := json.NewDecoder(e.Request().Body).Decode(&reqBody)
	if err1 != nil {
		return e.JSON(http.StatusBadRequest, enums.Faileddecode)
	}
	var v = validator.New()
	err2 := v.Struct(&reqBody) // checking validation
	if err2 != nil {
		return e.JSON(http.StatusBadRequest, enums.Validationerror)
	}

	// updating the user  instance provided
	err3 := service.Updateuser(reqBody, key)
	if err3 != nil {
		return e.JSON(http.StatusInternalServerError, enums.Servererror)
	}
	return e.JSON(http.StatusOK, enums.Successful)
}
