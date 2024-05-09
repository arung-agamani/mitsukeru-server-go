package services

import (
	"crypto/sha512"
	"errors"
	"github.com/arung-agamani/mitsukeru-server-go/db"
	"github.com/arung-agamani/mitsukeru-server-go/models"
	"github.com/arung-agamani/mitsukeru-server-go/utils/validator"
	"github.com/google/uuid"
)

type CreateUserPayload struct {
	Username    string `json:"username" validate:"required,alphanum,min=8,max=32"`
	DisplayName string `json:"displayName" validate:"max=64"`
	Bio         string `json:"bio" validate:"max=256"`
	Password    string `json:"password" validate:"required,min=8,max=32,alphanumunicode"`
}

type GetUserPayload struct {
	ID uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
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
	CreateUser(p CreateUserPayload) (*models.User, error)
	GetUser(p GetUserPayload) (*models.User, error)
	UpdateUserInfo(p UpdateUserInfoPayload) (*models.User, error)
	ChangeUserPassword(p ChangeUserPasswordPayload) (*models.User, error)
	DeleteUser(p DeleteUserPasswordPayload) (bool, error)
}

type userService struct{}

type ValidatorError struct {
	FieldErrors []validator.FieldError
}

type UserAlreadyExistError struct {
	ErrorMessage string
}

func (v ValidatorError) Error() string {
	return "Error on validation"
}

func (u UserAlreadyExistError) Error() string {
	return u.ErrorMessage
}

func NewUserService() UserService {
	return &userService{}
}

func (u *userService) CreateUser(p CreateUserPayload) (*models.User, error) {
	if ok, err := validator.ValidateStruct(&p); !ok {
		return nil, ValidatorError{FieldErrors: *err}
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
	createRes := dbConn.Create(&user)
	if ok, msg := db.HandleError(createRes.Error); !ok {
		return nil, errors.New(msg)
	}
	return &user, nil
}

func (u *userService) GetUser(p GetUserPayload) (*models.User, error) {
	if ok, err := validator.ValidateStruct(&p); !ok {
		return nil, ValidatorError{FieldErrors: *err}
	}
	dbConn := db.GetDB()
	user, err := db.GetUser(p.ID, dbConn)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userService) UpdateUserInfo(p UpdateUserInfoPayload) (*models.User, error) {
	if ok, err := validator.ValidateStruct(&p); !ok {
		return nil, ValidatorError{FieldErrors: *err}
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
	if ok, err := validator.ValidateStruct(&p); !ok {
		return nil, ValidatorError{FieldErrors: *err}
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
	if ok, err := validator.ValidateStruct(&p); !ok {
		return false, ValidatorError{FieldErrors: *err}
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
