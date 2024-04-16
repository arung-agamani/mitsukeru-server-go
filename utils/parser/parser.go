package parser

import (
	"encoding/json"
	"fmt"
	"github.com/arung-agamani/mitsukeru-go/responses"
	"github.com/arung-agamani/mitsukeru-go/utils/validator"
	"net/http"
)

func ParseStructPayload(w http.ResponseWriter, r *http.Request, target interface{}) bool {
	if valid, errMsg := validator.ValidateStruct(target); valid != true {
		responses.ErrResponse(w, &responses.ErrorResponse{
			Status:  400,
			Message: "Validation error",
			Error:   errMsg,
		})
		return false
	}
	return true
}

// ParseJSONPayload Parses request body as JSON request
// and stores it in target variable. This function
// is meant to be used in scope of HTTP request handler
// Parameters :
//
//	w http.ResponseWriter: Response writer for the current request
//	r *http.Request: Request object to the current request. Used to retrieve the request body
//	target interface{}: Target object that will have its fields filled with info fron JSON request
//
// Returns :
//
//	bool: indicating if the parsing success
func ParseJSONPayload(w http.ResponseWriter, r *http.Request, target interface{}) bool {
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&target); err != nil {
		var errMsg string
		if jsonErr, ok := err.(*json.UnmarshalTypeError); ok {
			errMsg = fmt.Sprintf("Error parsing. Field: %s. Reason: %s", jsonErr.Field, jsonErr.Error())
		} else if jsonErr, ok := err.(*json.SyntaxError); ok {
			errMsg = fmt.Sprintf("Error parsing. Request Body faults. Reason: %s", jsonErr.Error())
		} else {
			errMsg = "Invalid JSON format"
		}
		responses.ErrResponse(w, &responses.ErrorResponse{
			Status:  400,
			Message: "Invalid request body",
			Error:   errMsg,
		})
		return false
	}
	if valid, errMsg := validator.ValidateStruct(target); valid != true {
		responses.ErrResponse(w, &responses.ErrorResponse{
			Status:  400,
			Message: "Validation error",
			Error:   errMsg,
		})
		return false
	}
	return true
}
