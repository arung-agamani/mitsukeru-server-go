package services

import (
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"github.com/arung-agamani/mitsukeru-server-go/db"
	"github.com/arung-agamani/mitsukeru-server-go/models"
)

type AuthService interface {
	Login(username, password string) (*models.User, error)
	SignUp(username, password, confirmPassword string) (*models.User, error)
	Logout(username string) (bool, error)
}

type authService struct{}

func NewAuthService() AuthService { return &authService{} }

func (as *authService) Login(username, password string) (*models.User, error) {
	conn := db.GetDB()
	h := sha512.New()
	h.Write([]byte(password))
	bs := h.Sum(nil)
	user, err := db.GetUserByUsername(username, conn)
	if err != nil {
		return nil, err
	}
	if user.Password != hex.EncodeToString(bs) {
		return nil, errors.New("wrong password")
	}
	return user, nil
}

func (as *authService) SignUp(username, password, confirmPassword string) (*models.User, error) {
	if password != confirmPassword {
		return nil, errors.New("field 'password' and 'confirmPassword' have mismatch values")
	}
	conn := db.GetDB()
	if user, _ := db.GetUserByUsername(username, conn); user != nil {
		return nil, UserAlreadyExistError{ErrorMessage: "user already exist ya"}
	}
	h := sha512.New()
	h.Write([]byte(password))
	bs := h.Sum(nil)
	user := models.User{
		Username:    username,
		DisplayName: "",
		Bio:         "",
		Password:    hex.EncodeToString(bs),
	}
	create := conn.Create(&user)
	if ok, msg := db.HandleError(create.Error); !ok {
		return nil, errors.New(msg)
	}
	return &user, nil
}

func (as *authService) Logout(username string) (bool, error) {
	return true, nil
}
