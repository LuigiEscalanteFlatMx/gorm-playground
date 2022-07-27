package port

import "github.com/luigiescalante/proyect-template/core/domain"

type UsersRepository interface {
	Save(user *domain.Users) error
	Delete(user *domain.Users) error
	GetById(id int) (user *domain.Users, err error)
	GetUsers(page int, totalRecords int) (map[string]interface{}, error)
}
