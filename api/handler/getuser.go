package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"

	"awesomeProject3/api/service"
	"awesomeProject3/pkg/enums"
)

func GetUser(e echo.Context) error {
	Id, err := strconv.Atoi(e.QueryParam("id"))

	if err != nil {
		return e.JSON(http.StatusBadRequest, enums.Paramserror)
	}
	// this function helps to find the user with the particular id
	ans, err := service.GetUser(bson.M{"Id": Id}, e.QueryParam("id"))
	if err != nil {
		return e.JSON(http.StatusInternalServerError, enums.Servererror)
	}
	return e.JSON(http.StatusOK, ans)
}
