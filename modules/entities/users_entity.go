package entities

type UsersRegisterReq struct {
	FacebookId string `bson:"facebook_id" json:"facebook_id"`
	Email      string `bson:"email" json:"email"`
}

type UsersCredential struct {
	FacebookId string `bson:"facebook_id" json:"facebook_id"`
	Email      string `bson:"email" json:"email"`
}
