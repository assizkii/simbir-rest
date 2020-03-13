package interfaces

import "github.com/assizkii/simbir-rest/internal/entities"

type AppStorage interface {
	Validate(account entities.User) error
	Get(login string) (*entities.User, error)
	Add(account entities.User) error
}
