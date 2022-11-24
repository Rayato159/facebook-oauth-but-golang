package entities

type UsersRegisterReq struct {
	Id    string `bson:"_id" json:"id"`
	Email string `bson:"email" json:"email"`
}

type UsersCredential struct {
	Id          string `bson:"_id" json:"id"`
	Email       string `bson:"email" json:"email"`
	AccessToken string `bson:"access_token" json:"access_token"`
}
