package controllers

import (
	"github.com/arung-agamani/mitsukeru-server-go/config"
	"github.com/arung-agamani/mitsukeru-server-go/responses"
	"github.com/arung-agamani/mitsukeru-server-go/services"
	"github.com/arung-agamani/mitsukeru-server-go/utils/logger"
	"github.com/arung-agamani/mitsukeru-server-go/utils/parser"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"net/http"
	"time"
)

// swagger:parameters LoginPayload
type LoginPayload struct {
	Username string `json:"username" validate:"required,min=8,max=32"`
	Password string `json:"password" validate:"required,min=8,max=32"`
}

// swagger:response LoginResponse
type LoginResponse struct {
	ID          uuid.UUID `json:"id"`
	Username    string    `json:"username"`
	DisplayName string    `json:"displayName"`
	Bio         string    `json:"bio"`
}

// swagger:parameters SigninPayload
type SigninPayload struct {
	Username        string `json:"username" validate:"required,min=8,max=32"`
	Password        string `json:"password" validate:"required,min=8,max=32"`
	ConfirmPassword string `json:"confirmPassword" validate:"required,eqfield=Password"`
}

// swagger:route POST /auth/login LoginPayload LoginResponse
// Login to Mitsukeru
//
// ---
// responses:
//
//	200: LoginResponse
func LoginHandler(deps services.Dependencies) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var payload LoginPayload
		if parser.ParseJSONPayload(w, r, &payload) != true {
			return
		}
		user, err := deps.AuthService.Login(payload.Username, payload.Password)
		if err != nil {
			logger.Error(err)
			responses.HandleError(w, err)
			return
		}
		t := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
			"username":    user.Username,
			"displayName": user.DisplayName,
			"bio":         user.Bio,
		})
		s, err := t.SignedString([]byte(config.GetJwtSecret()))
		if err != nil {
			logger.Error(err)
			responses.ErrResponse(w, &responses.ErrorResponse{
				Status:  500,
				Message: "Error when generating JWT signed string",
				Error:   nil,
			})
			return
		}
		http.SetCookie(w, &http.Cookie{
			Name:    "ninshin-pompu",
			Value:   s,
			Expires: time.Now().Add(365 * 24 * time.Hour),
		})
		resBody := LoginResponse{
			ID:          user.ID,
			Username:    user.Username,
			DisplayName: user.DisplayName,
			Bio:         user.Bio,
		}
		responses.OkResponse(w, &responses.Response{
			Status:  200,
			Message: "Authenticated",
			Data:    resBody,
		})
	}
}

func SignInHandler(deps services.Dependencies) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var payload SigninPayload
		if parser.ParseJSONPayload(w, r, &payload) != true {
			return
		}
		user, err := deps.AuthService.SignUp(payload.Username, payload.Password, payload.ConfirmPassword)
		if err != nil {
			responses.ErrResponse(w, &responses.ErrorResponse{
				Status:  400,
				Message: "Bad request",
				Error:   err.Error(),
			})
			return
		}
		t := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
			"username":    user.Username,
			"displayName": user.DisplayName,
			"bio":         user.Bio,
		})
		s, err := t.SignedString([]byte(config.GetJwtSecret()))
		if err != nil {
			logger.Error(err)
			responses.ErrResponse(w, &responses.ErrorResponse{
				Status:  500,
				Message: "Error when generating JWT signed string",
				Error:   nil,
			})
			return
		}
		http.SetCookie(w, &http.Cookie{
			Name:    "ninshin-pompu",
			Value:   s,
			Expires: time.Now().Add(365 * 24 * time.Hour),
		})

		responses.OkResponse(w, &responses.Response{
			Status:  201,
			Message: "User created",
			Data:    user,
		})
	}
}

func LogoutHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.SetCookie(w, &http.Cookie{
			Name:    "ninshin-pompu",
			Value:   "",
			Expires: time.Now(),
		})
		responses.OkResponse(w, &responses.Response{
			Status:  200,
			Message: "logged out",
			Data:    nil,
		})
	}
}
