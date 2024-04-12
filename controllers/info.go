package controllers

import (
	"github.com/arung-agamani/mitsukeru-go/config"
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
		res := &Response{
			Status:  200,
			Message: "Info",
			Data:    infoBody,
		}
		OkResponse(w, res)
	}
}
