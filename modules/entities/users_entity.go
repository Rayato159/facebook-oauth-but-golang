package entities

type UsersContext string

const (
	UsersCon UsersContext = "UsersController"
	UsersUse UsersContext = "UsersUsecase"
	UsersRep UsersContext = "UsersRepository"
)

type UsersRepository interface{}
type UsersUsecase interface{}

type UsersCredential struct {
	UserId      string `bson:"user_id" json:"user_id"`
	Name        string `bson:"name" json:"name"`
	AccessToken string `bson:"access_token" json:"access_token"`
}
