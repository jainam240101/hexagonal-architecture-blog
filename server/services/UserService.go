package users

import (
	"github.com/google/uuid"
	users "github.com/jainam240101/todo/server/domain/Users"
	"golang.org/x/crypto/bcrypt"
)

type DefaultUserService struct {
	repo users.UserRepositoryDb
}

type UserService interface {
	CreateUser(users.UserModel) (*users.UserModel, error)
	FindUserByUsername(string) (*users.UserModel, error)
	UpdateUser(users.UserModel,string) (*users.UserModel, error)
	DeleteUser(string) error
}

func NewCustomerService(repository users.UserRepositoryDb) DefaultUserService {
	return DefaultUserService{repo: repository}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 4)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func (db DefaultUserService) CreateUser(u users.UserModel) (*users.UserModel, error) {
	pass, err := HashPassword(u.Password)
	if err != nil {
		return nil, err
	}
	u.Password = pass
	u.ID = uuid.New()
	user, err := db.repo.CreateUser(u)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (db DefaultUserService) FindUserByUsername(username string) (*users.UserModel, error) {
	user, err := db.repo.FindUserById(username)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (db DefaultUserService) UpdateUser(u users.UserModel, userid string) (*users.UserModel, error) {
	if u.Password != "" {
		pass, err := HashPassword(u.Password)
		u.Password = pass
		if err != nil {
			return nil, err
		}
	}
	dbValue, err := db.repo.UpdateUser(userid, u)
	if err != nil {
		return nil, err
	}
	return dbValue, nil
}

func (db DefaultUserService) DeleteUser(userId string) error {
	err := db.repo.DeleteUser(userId)
	if err != nil {
		return err
	}
	return nil
}
