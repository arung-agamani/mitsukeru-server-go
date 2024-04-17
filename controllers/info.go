package controllers

import (
	"github.com/arung-agamani/mitsukeru-server-go/config"
	"github.com/arung-agamani/mitsukeru-server-go/responses"
	"net/http"
)

type Info struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

func InfoHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		infoBody := Info{
			Name:    config.GetAppName(),
			Version: config.GetVersion(),
		}
		res := &responses.Response{
			Status:  200,
			Message: "Info",
			Data:    infoBody,
		}
		responses.OkResponse(w, res)
	}
}
