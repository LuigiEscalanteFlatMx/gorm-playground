package domain

import (
	"gorm.io/gorm"
	"time"
)

type Roles struct {
	Id        int            `json:"id,primary_key"`
	UserId    int            `json:"user_id"`
	User      Users          `gorm:"foreignKey:UserId"`
	Role      string         `json:"role" gorm:"size:50"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"default:null"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
