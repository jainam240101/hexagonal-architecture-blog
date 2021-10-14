package users

import (
	"gorm.io/gorm"
)

type UserRepositoryDb struct {
	Client *gorm.DB
}

func NewUserRepositoryDb(client *gorm.DB) UserRepositoryDb {
	return UserRepositoryDb{Client: client}
}

func (d UserRepositoryDb) CreateUser(u UserModel) (*UserModel, error) {
	if err := d.Client.Save(&u).Error; err != nil {
		return nil,err
	}
	return &u, nil
}
