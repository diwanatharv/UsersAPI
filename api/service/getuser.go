package service

import (
	"encoding/json"
	"fmt"

	"github.com/redis/go-redis/v9"

	Mongo "awesomeProject3/pkg/dataaccess/mongo"
	Redis "awesomeProject3/pkg/dataaccess/redis"
	"awesomeProject3/pkg/models"
)

func GetUser(filter interface{}, key string) (models.User, error) {
	var db = Mongo.Mongosmanager()
	var cache = Redis.RedisManager()
	var ans models.User
	val, err := cache.Get(key).Result()
	if err == redis.Nil {
		var findOneUser models.User
		// finding in  the database
		err = db.FindOne(filter).Decode(&findOneUser)

		if err != nil {
			return ans, err
		}
		byteans, err := json.Marshal(findOneUser) // converting to byte for setting in the redis
		if err != nil {
			return ans, err
		}
		err3 := cache.Set(key, byteans, 0).Err()
		if err3 != nil {
			return ans, err3
		}
		return findOneUser, nil
	} else {
		err := json.Unmarshal([]byte(val), &ans)
		if err != nil {
			fmt.Println("error in unmarshalling")
		}
		return ans, nil
	}

}
