package responses

import (
	"encoding/json"
	"errors"
	"github.com/arung-agamani/mitsukeru-server-go/services"
	"github.com/arung-agamani/mitsukeru-server-go/utils/logger"
	"net/http"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
}

func OkResponse(w http.ResponseWriter, response *Response) {
	w.Header().Set("Content-Type", "application/json")
	resJson, err := json.Marshal(response)
	if err != nil {
		logger.Errorf("Unable to marshall JSON request")
	}
	w.WriteHeader(response.Status)
	_, _ = w.Write(resJson)
}

func ErrResponse(w http.ResponseWriter, response *ErrorResponse) {
	w.Header().Set("Content-Type", "application/json")
	resJson, err := json.Marshal(response)
	if err != nil {
		logger.Errorf("Unable to marshall JSON request")
	}
	w.WriteHeader(response.Status)
	_, _ = w.Write(resJson)
}

func HandleError(w http.ResponseWriter, err error) {
	if errors.Is(err, services.ValidatorError{}) {
		ErrResponse(w, &ErrorResponse{
			Status:  400,
			Message: "Bad request",
			Error:   err,
		})
		return
	}
	ErrResponse(w, &ErrorResponse{
		Status:  500,
		Message: "Internal server error",
		Error:   err,
	})
	return
}
