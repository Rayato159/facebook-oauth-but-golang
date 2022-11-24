package repositories

import (
	"context"
	"fmt"
	"time"

	"github.com/Rayato159/facebook-oauth-but-go/modules/entities"
	"github.com/Rayato159/facebook-oauth-but-go/package/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type facebookRepository struct {
	Db *mongo.Database
}

func NewFacebookRepository(db *mongo.Database) *facebookRepository {
	return &facebookRepository{
		Db: db,
	}
}

func (r *facebookRepository) InsertOneUsers(ctx context.Context, req *entities.UsersRegisterReq) (*entities.UsersCredential, error) {
	defer fmt.Printf("return:\t%v.%v\n", entities.FacebookRep, utils.GetFunctionName())
	ctx = context.WithValue(ctx, entities.FacebookRep, time.Now().UnixMicro())
	fmt.Printf("called:\t%v.%v\n", entities.FacebookRep, utils.GetFunctionName())

	coll := r.Db.Collection("users")
	result, err := coll.InsertOne(ctx, req)
	if err != nil {
		return nil, err
	}
	user := new(entities.UsersCredential)
	userResult := coll.FindOne(ctx, bson.D{{"_id", result.InsertedID}})
	if err := userResult.Decode(&user); err != nil {
		return nil, fmt.Errorf("error, can't decode user with an error: %v", err.Error())
	}
	return user, nil
}
