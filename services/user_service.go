package services

import (
	"github.com/kataras/iris/v12/x/errors"
	"go-product/datamodels"
	"go-product/repositories"
	"golang.org/x/crypto/bcrypt"
)

type IUserService interface {
	// check if account and password are matching
	IsPwdSuccess(userName string, pwd string) (user *datamodels.User, isOk bool)
	AddUser(user *datamodels.User) (userId int64, err error)
}

type UserService struct {
	UserRepository repositories.IUserRepository
}

func NewService(repository repositories.IUserRepository) IUserService {
	return &UserService{repository}
}

// self-defined function
func ValidatePassword(userPassword string, hashed string) (isOK bool, err error) {
	if err = bcrypt.CompareHashAndPassword([]byte(hashed), []byte(userPassword)); err != nil {
		return false, errors.New("Password is incorrect!")
	}
	return true, nil
}

func (u *UserService) IsPwdSuccess(userName string, pwd string) (user *datamodels.User, isOk bool) {
	// get the user struct first
	user, err := u.UserRepository.Select(userName)
	if err != nil {
		return
	}
	// check is password is correct
	isOk, _ = ValidatePassword(pwd, user.HashPassword)

	if !isOk {
		return &datamodels.User{}, false
	}
	return
}

func GeneratePassword(userPassword string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.DefaultCost)
}

func (u *UserService) AddUser(user *datamodels.User) (userId int64, err error) {
	pwdByte, errPwd := GeneratePassword(user.HashPassword)
	if errPwd != nil {
		return userId, errPwd
	}

	user.HashPassword = string(pwdByte)
	return u.UserRepository.Insert(user)
}
