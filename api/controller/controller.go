package controller

import (
	"errors"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"

	"awesomeProject3/api/handler"
)

func viperEnvVariable(key string) string {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}
	value, ok := viper.Get(key).(string)
	if !ok {
		log.Fatalf("Invalid type assertion")
	}

	return value
}
func basicAuthValidator(username string, password string, c echo.Context) (bool, error) {
	viperenv1 := viperEnvVariable("USERNAME")
	viperenv2 := viperEnvVariable("PASSWORD")

	if username == viperenv1 && password == viperenv2 {
		return true, nil
	} else {
		return false, errors.New("password and username do not match")
	}
}
func Createroutes(e *echo.Echo) {
	e.Use(middleware.BasicAuth(basicAuthValidator))
	e.POST("user", handler.CreateUser)
	e.GET("users", handler.GetallUsers)
	e.GET("user", handler.GetUser)
	e.PUT("user", handler.UpdateUser)

}
