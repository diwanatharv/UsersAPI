package mongo

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"awesomeProject3/pkg/config"
	"awesomeProject3/pkg/models"
)

type mongomethods interface {
	FindOne(interface{}) *mongo.SingleResult
	FindAll(context.Context, interface{}, ...*options.FindOptions) (*mongo.Cursor, error)
	InsertOne(user models.User) (*mongo.InsertOneResult, error)
	UpdateOne(interface{}, interface{}) (*mongo.UpdateResult, error)
	TotalDocument() (int64, error)
	FindOneAndUpdate(ctx context.Context, filter interface{}, update interface{}, opts *options.FindOneAndUpdateOptions) *mongo.SingleResult
}
type Mongos struct {
	users *mongo.Collection
}

func Mongosmanager() *Mongos {
	return &Mongos{
		users: config.CreateMongoDatabase("Userdb").Collection("user"),
	}
}
func (db *Mongos) FindOne(filter interface{}) *mongo.SingleResult {

	return db.users.FindOne(context.Background(), filter)
}

func (db *Mongos) FindAll(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error) {
	return db.users.Find(ctx, filter)
}

func (db *Mongos) InsertOne(reqBody models.User) (*mongo.InsertOneResult, error) {
	return db.users.InsertOne(context.Background(), reqBody)
}

func (db *Mongos) UpdateOne(filter interface{}, update interface{}) (*mongo.UpdateResult, error) {
	return db.users.UpdateOne(context.Background(), filter, update)
}

func (db *Mongos) TotalDocument() (int64, error) {
	if db.users == nil {
		fmt.Println("there is no document here")
	}
	len, err := db.users.EstimatedDocumentCount(context.Background())
	return len, err
}
func (db *Mongos) FindOneAndUpdate(ctx context.Context, filter interface{}, update interface{}, opts *options.FindOneAndUpdateOptions) *mongo.SingleResult {
	return db.users.FindOneAndUpdate(ctx, filter, update, opts)
}
