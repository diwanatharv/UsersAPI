package service

import (
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/mongo"

	Mongo "awesomeProject3/pkg/dataaccess/mongo"

	"awesomeProject3/pkg/models"
)

func Createuser(reqbody models.User) (*mongo.InsertOneResult, error) {
	db := Mongo.Mongosmanager()
	ans, err := db.TotalDocument()
	if err != nil {
		log.Error("error in finding total document")
		return nil, err
	}
	// users id equal to the lenght of the reqbody
	reqbody.Id = int(ans)
	res, err := db.InsertOne(reqbody)
	if err != nil {
		log.Error("error in insertion in mongo")
		return nil, err
	}
	return res, nil
}
