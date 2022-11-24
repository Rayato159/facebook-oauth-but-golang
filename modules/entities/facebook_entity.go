package entities

import "context"

type FacebookContext string

const (
	FacebookCon FacebookContext = "FacebookController"
	FacebookUse FacebookContext = "FacebookUsecase"
	FacebookRep FacebookContext = "FacebookRepository"
	FacebookExt FacebookContext = "FacebookContext"
)

type FacebookUsecase interface {
	InsertOneUsers(ctx context.Context, req *UsersRegisterReq) (*UsersCredential, error)
}
type FacebookRepository interface {
	InsertOneUsers(ctx context.Context, req *UsersRegisterReq) (*UsersCredential, error)
}
type FacebookApi interface{}
