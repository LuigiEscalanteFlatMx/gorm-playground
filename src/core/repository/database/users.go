package database

import (
	"github.com/luigiescalante/proyect-template/core/domain"
	"github.com/luigiescalante/proyect-template/infrastructure/repository/database"
)

type repositoryUsers struct {
}

func (repo repositoryUsers) GetUsers(page int, totalRecords int) (map[string]interface{}, error) {
	var totalRows int64
	var rows []interface{}
	db := database.Engine()
	result := make(map[string]interface{})
	db.Model(&domain.Users{}).Count(&totalRows)
	rowsData, _ := db.Model(&domain.Users{}).Limit(totalRecords).Offset(GetOffset(page, totalRecords)).Order("id desc").Rows()
	for rowsData.Next() {
		user := domain.Users{}
		db.ScanRows(rowsData, &user)
		rows = append(rows, user)
	}
	result["totalRecords"] = totalRows
	result["records"] = rows
	return result, nil
}

func (repo repositoryUsers) GetById(id int) (*domain.Users, error) {
	var user *domain.Users
	tx := database.Engine().Where("id=?", id).First(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return user, nil
}

func UserRepositoryFactory() *repositoryUsers {
	return &repositoryUsers{}
}

func (repo repositoryUsers) Save(user *domain.Users) error {
	tx := database.Engine().Save(user)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (repo repositoryUsers) Delete(user *domain.Users) error {
	tx := database.Engine().Delete(user)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
