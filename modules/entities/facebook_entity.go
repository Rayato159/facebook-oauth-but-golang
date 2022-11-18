package entities

type FacebookContext string

const (
	FacebookCon FacebookContext = "FacebookController"
	FacebookUse FacebookContext = "FacebookUsecase"
	FacebookRep FacebookContext = "FacebookRepository"
	FacebookExt FacebookContext = "FacebookContext"
)

type FacebookController interface{}
type FacebookUsecase interface{}
type FacebookRepository interface{}
type FacebookApi interface{}
