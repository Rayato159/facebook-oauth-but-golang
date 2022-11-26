package entities

type FacebookContext string

const (
	FacebookCon FacebookContext = "FacebookController"
	FacebookUse FacebookContext = "FacebookUsecase"
	FacebookRep FacebookContext = "FacebookRepository"
	FacebookExt FacebookContext = "FacebookContext"
)

type FacebookUsecase interface{}
type FacebookApi interface{}

type AccessTokenRes struct {
	AccessToken string `bson:"access_token" json:"access_token"`
	TokenType   string `bson:"token_type" json:"token_type"`
	ExpiresIn   int    `bson:"expires_in" json:"expires_in"`
}

type FacebookGraphUsersDataRes struct {
	UserId string `bson:"user_id" json:"id"`
	Email  string `bson:"email" json:"email"`
	Name   string `bson:"name" json:"name"`
}
