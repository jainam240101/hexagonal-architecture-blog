package users

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:char(36);primary_key"`
	Name     string    `json:"name"`
	Email    string    `json:"email" gorm:"uniqueIndex"`
	Password string    `json:"password"`
	Username string    `json:"username" gorm:"uniqueIndex"`
}

type UserRepository interface {
	CreateUser(UserModel) (*UserModel, error)
}
