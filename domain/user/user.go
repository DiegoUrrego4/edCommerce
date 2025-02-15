package user

import "github.com/DiegoUrrego4/edCommerce/model"

// UseCase -> Puerto o servicio
type UseCase interface {
	Create(newUser *model.User) error
	GetByEmail(email string) (model.User, error)
	GetAll() (model.Users, error)
}

// Storage -> Adaptador o Repository
type Storage interface {
	Create(newUser *model.User) error
	GetByEmail(email string) (model.User, error)
	GetAll() (model.Users, error)
}
