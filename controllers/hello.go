package controllers

import "net/http"

func HelloHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res := &Response{
			Status:  200,
			Message: "Awoo!",
			Data:    nil,
		}
		OkResponse(w, res)
	}
}
