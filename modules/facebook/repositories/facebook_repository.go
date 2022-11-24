package repositories

import (
	"context"
	"fmt"
	"time"

	"github.com/Rayato159/facebook-oauth-but-go/modules/entities"
	"github.com/Rayato159/facebook-oauth-but-go/package/utils"
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
	ctx = context.WithValue(ctx, entities.FacebookCon, time.Now().UnixMicro())
	fmt.Printf("called:\t%v.%v\n", entities.FacebookRep, utils.GetFunctionName())

	coll := r.Db.Collection("users")
	if _, err := coll.InsertOne(ctx, req); err != nil {
		return nil, err
	}
	return nil, nil
}
