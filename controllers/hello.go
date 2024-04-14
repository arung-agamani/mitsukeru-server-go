package controllers

import (
	"github.com/arung-agamani/mitsukeru-go/responses"
	"net/http"
)

func HelloHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res := &responses.Response{
			Status:  200,
			Message: "Awoo!",
			Data:    nil,
		}
		responses.OkResponse(w, res)
	}
}
