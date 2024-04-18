package interfaces

import (
	"github.com/jfelipearaujo-org/lambda-register/src/entities"
)

type Database interface {
	CheckIfCPFIsInUse(cpf string) (bool, error)
	PersistUser(user entities.User) error
}
