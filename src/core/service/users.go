package service

import (
	"github.com/luigiescalante/proyect-template/core/domain"
	"github.com/luigiescalante/proyect-template/core/port"
)

type userService struct {
	User       *domain.Users
	Repository port.UsersRepository
}

func UserServiceFactory(user *domain.Users, repo port.UsersRepository) *userService {
	return &userService{
		User:       user,
		Repository: repo,
	}
}

func (service userService) Save() error {
	err := service.Repository.Save(service.User)
	if err != nil {
		return err
	}
	return nil
}

func (service userService) Delete() error {
	err := service.Repository.Delete(service.User)
	if err != nil {
		return err
	}
	return nil
}

func (service userService) GetById(id int) (*domain.Users, error) {
	user, err := service.Repository.GetById(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (service userService) GetUsers(page int, totalRecords int) (map[string]interface{}, error) {
	userRows, err := service.Repository.GetUsers(page, totalRecords)
	if err != nil {
		return nil, err
	}
	return userRows, nil
}
