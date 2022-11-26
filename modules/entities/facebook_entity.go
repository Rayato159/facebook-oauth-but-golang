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
