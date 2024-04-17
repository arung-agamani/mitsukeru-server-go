package controllers

import (
	"errors"
	"github.com/arung-agamani/mitsukeru-server-go/responses"
	"github.com/arung-agamani/mitsukeru-server-go/services"
	"github.com/arung-agamani/mitsukeru-server-go/utils/parser"
	"net/http"
)

type CreateUserRequestPayload struct {
	Username        string `json:"username" validate:"required,alphanum,min=8,max=32"`
	DisplayName     string `json:"displayName" validate:"max=64"`
	Bio             string `json:"bio" validate:"max=256"`
	Password        string `json:"password" validate:"required,min=8,max=32,alphanumunicode"`
	ConfirmPassword string `json:"confirmPassword" validate:"required,min=8,max=32,alphanumunicode,eqfield=password"`
}

func CreateUserHandler(deps services.Dependencies) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var payload CreateUserRequestPayload
		if parser.ParseJSONPayload(w, r, &payload) != true {
			return
		}
		user, err := deps.UserService.CreateUser(services.CreateUserPayload{
			Username:    payload.Username,
			DisplayName: payload.DisplayName,
			Bio:         payload.Bio,
			Password:    payload.Password,
		})
		if err != nil {
			if errors.Is(err, services.ValidatorError{}) {
				responses.ErrResponse(w, &responses.ErrorResponse{
					Status:  400,
					Message: "Bad request",
					Error:   err,
				})
			}
			responses.ErrResponse(w, &responses.ErrorResponse{
				Status:  500,
				Message: "Internal server error",
				Error:   err,
			})
			return
		}
		responses.OkResponse(w, &responses.Response{
			Status:  201,
			Message: "User has been created",
			Data:    user,
		})
	}
}
func GetUserHandler(deps services.Dependencies) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
func UpdateUserInfoHandler(deps services.Dependencies) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
func ChangeUserPasswordHandler(deps services.Dependencies) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
func DeleteUserHandler(deps services.Dependencies) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
