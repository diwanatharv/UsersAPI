package service

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	Mongo "awesomeProject3/pkg/dataaccess/mongo"
	Redis "awesomeProject3/pkg/dataaccess/redis"
	"awesomeProject3/pkg/models"
)

func Updateuser(reqBody models.User, key string) error {
	var db = Mongo.Mongosmanager()
	var cache = Redis.RedisManager()
	keyInt, _ := strconv.Atoi(key)
	filter := bson.M{"Id": keyInt}

	update := bson.M{"$set": bson.M{"FirstName": reqBody.FirstName, "LastName": reqBody.LastName, "Company": reqBody.Company, "PhoneNumber": reqBody.PhoneNumber, "Email": reqBody.Email, "Country": reqBody.Country, "BuisnessType": reqBody.BuisnessType}}

	opts := options.FindOneAndUpdate().SetUpsert(true)
	reqBody.Id = keyInt
	// Perform the upsert
	var result bson.M
	err := db.FindOneAndUpdate(context.Background(), filter, update, opts).Decode(&result)
	if err != nil {
		return err
	}
	reqBodyv2, err := json.Marshal(reqBody)
	if err != nil {
		return err
	}

	err = cache.Set(key, reqBodyv2, 0).Err()

	if err == redis.Nil {
		return err
	}
	return nil
}
