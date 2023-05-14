package service

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"

	Mongo "awesomeProject3/pkg/dataaccess/mongo"
	"awesomeProject3/pkg/models"
)

func Getalluser(allUser []models.User) ([]models.User, error) {
	db := Mongo.Mongosmanager()
	filter := bson.M{
		"Id": bson.M{"$exists": true},
	}
	//function for finding all the user
	var oneUser models.User
	//will find all the id through filter condition
	findElementRes, err := db.FindAll(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	//iterating through the slice
	for findElementRes.Next(context.Background()) {
		err := findElementRes.Decode(&oneUser)
		if err != nil {
			fmt.Println(err)
		}
		allUser = append(allUser, oneUser)
	}
	return allUser, nil

}
