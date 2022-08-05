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
	offset := GetOffset(page, totalRecords)
	db.Model(&domain.Users{}).Count(&totalRows)
	/*Query from Model*/
	rowsData, _ := db.Debug().Model(&domain.Users{}).Limit(totalRecords).Offset(offset).Order("id desc").Rows()
	/*Query Builder from methods */
	/*querySemiManual := db.Debug().Select("*").Table("users")
	if offset <= 0 {
		querySemiManual.Where("users.deleted_at IS NULL ORDER BY id desc LIMIT ?", totalRecords)
	} else {
		querySemiManual.Where("users.deleted_at IS NULL ORDER BY id desc LIMIT ? OFFSET ?", totalRecords, offset)
	}
	rowsData, _ := querySemiManual.Rows()*/
	/*Query direct*/
	/*var rowsData *sql.Rows
	if offset <= 0 {
		rowsData, _ = db.Debug().Raw("SELECT * FROM users WHERE users.deleted_at IS NULL ORDER BY id desc LIMIT ?", totalRecords).Rows()
	} else {
		rowsData, _ = db.Debug().Raw("SELECT * FROM users WHERE users.deleted_at IS NULL ORDER BY id desc LIMIT ? OFFSET ?", totalRecords, offset).Rows()
	}*/
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
	db := database.Engine()
	tx := db.Where("id=?", id).First(&user)
	//Users with join
	//tx := db.Debug().Joins("inner join roles on roles.user_id=users.id").Where("users.id=?", id).First(&user)
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
