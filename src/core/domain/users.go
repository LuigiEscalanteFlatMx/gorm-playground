package domain

import (
	"gorm.io/gorm"
	"time"
)

type Users struct {
	Id        int            `json:"id,primary_key"`
	FirstName string         `json:"first_name" gorm:"size:50"`
	LastName  string         `json:"last_name" gorm:"size:50"`
	Email     string         `json:"email" gorm:"size:50;unique"`
	IsAdult   bool           `json:"is_adult" gorm:"default:true"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"default:null"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
