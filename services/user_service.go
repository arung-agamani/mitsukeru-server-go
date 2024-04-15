package services

import (
	"crypto/sha512"
	"errors"
	"github.com/arung-agamani/mitsukeru-go/db"
	"github.com/arung-agamani/mitsukeru-go/models"
	"github.com/arung-agamani/mitsukeru-go/utils/validator"
	"github.com/google/uuid"
)

type CreateUserPayload struct {
	Username    string `json:"username" validate:"required,alphanum,min=8,max=32"`
	DisplayName string `json:"displayName" validate:"max=64"`
	Bio         string `json:"bio" validate:"max=256"`
	Password    string `json:"password" validate:"required,min=8,max=32,alphanumunicode"`
}

type UpdateUserInfoPayload struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	DisplayName string    `json:"displayName" validate:"max=64"`
	Bio         string    `json:"bio" validate:"max=256"`
}

type ChangeUserPasswordPayload struct {
	ID       uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	Password string    `json:"password" validate:"required,min=8,max=32,alphanumunicode"`
}

type DeleteUserPasswordPayload struct {
	ID uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
}

type UserService interface {
	CreateUser(p CreateUserPayload) (bool, error)
	UpdateUserInfo(p UpdateUserInfoPayload) (*models.User, error)
	ChangeUserPassword(p ChangeUserPasswordPayload) (*models.User, error)
	DeleteUser(p DeleteUserPasswordPayload) (bool, error)
}

type userService struct{}

func NewUserService() UserService {
	return &userService{}
}

func (u *userService) CreateUser(p CreateUserPayload) (bool, error) {
	if ok, msg := validator.ValidateStruct(&p); !ok {
		return false, errors.New(msg)
	}
	dbConn := db.GetDB()
	h := sha512.New()
	h.Write([]byte(p.Password))
	bs := h.Sum(nil)
	user := models.User{
		Username:    p.Username,
		DisplayName: p.DisplayName,
		Bio:         p.Bio,
		Password:    string(bs),
	}
	createRes := dbConn.Create(user)
	if ok, msg := db.HandleError(createRes.Error); !ok {
		return ok, errors.New(msg)
	}
	return true, nil
}
func (u *userService) UpdateUserInfo(p UpdateUserInfoPayload) (*models.User, error) {
	if ok, msg := validator.ValidateStruct(&p); !ok {
		return nil, errors.New(msg)
	}
	dbConn := db.GetDB()
	user, err := db.GetUser(p.ID, dbConn)
	if err != nil {
		return nil, err
	}
	user.DisplayName = p.DisplayName
	user.Bio = p.Bio
	res := dbConn.Save(&user)
	if ok, err := db.HandleError(res.Error); !ok {
		return user, errors.New(err)
	}
	return user, nil
}
func (u *userService) ChangeUserPassword(p ChangeUserPasswordPayload) (*models.User, error) {
	if ok, msg := validator.ValidateStruct(&p); !ok {
		return nil, errors.New(msg)
	}
	dbConn := db.GetDB()
	user, err := db.GetUser(p.ID, dbConn)
	if err != nil {
		return nil, err
	}
	h := sha512.New()
	h.Write([]byte(p.Password))
	bs := h.Sum(nil)
	user.Password = string(bs)
	res := dbConn.Save(&user)
	if ok, err := db.HandleError(res.Error); !ok {
		return nil, errors.New(err)
	}
	return user, nil
}
func (u *userService) DeleteUser(p DeleteUserPasswordPayload) (bool, error) {
	if ok, msg := validator.ValidateStruct(&p); !ok {
		return false, errors.New(msg)
	}
	dbConn := db.GetDB()
	user, err := db.GetUser(p.ID, dbConn)
	if err != nil {
		return false, err
	}
	res := dbConn.Delete(&user)
	if ok, err := db.HandleError(res.Error); !ok {
		return false, errors.New(err)
	}
	return true, nil
}
