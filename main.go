// Package classification awesome.
//
// Documentation of our awesome API.
//
//	 Schemes: http
//	 BasePath: /
//	 Version: 1.0.0
//	 Host: some-url.com
//
//	 Consumes:
//	 - application/json
//
//	 Produces:
//	 - application/json
//
//	 Security:
//	 - basic
//
//	SecurityDefinitions:
//	basic:
//	  type: basic
//
// swagger:meta
package main

import (
	"context"
	"errors"
	"github.com/arung-agamani/mitsukeru-server-go/config"
	"github.com/arung-agamani/mitsukeru-server-go/controllers"
	"github.com/arung-agamani/mitsukeru-server-go/db"
	"github.com/arung-agamani/mitsukeru-server-go/services"
	"github.com/arung-agamani/mitsukeru-server-go/utils/logger"
	"github.com/arung-agamani/mitsukeru-server-go/utils/validator"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	config.InitConfig()
	logger.InitLogger()
	db.InitDb()
	db.InitS3Client()
	validator.Init()

	deps := services.NewDependencies()
	router := mux.NewRouter()
	router.HandleFunc("/", controllers.HelloHandler()).Methods(http.MethodGet)
	router.HandleFunc("/info", controllers.InfoHandler()).Methods(http.MethodGet)
	router.HandleFunc("/lost-item", controllers.CreateLostItemHandler(deps)).Methods(http.MethodPost)
	router.HandleFunc("/lost-item", controllers.ListLostItemHandler(deps)).Methods(http.MethodGet)
	router.HandleFunc("/lost-items/{itemId}", controllers.GetLostItemHandler(deps)).Methods(http.MethodGet)
	router.HandleFunc("/lost-items/{itemId}", controllers.UpdateLostItemHandler(deps)).Methods(http.MethodPost)
	router.HandleFunc("/lost-items/{itemId}", controllers.DeleteLostItemHandler(deps)).Methods(http.MethodDelete)

	router.HandleFunc("/auth/login", controllers.LoginHandler(deps)).Methods(http.MethodPost)
	router.HandleFunc("/auth/signup", controllers.SignInHandler(deps)).Methods(http.MethodPost)
	router.HandleFunc("/auth/logout", controllers.LogoutHandler()).Methods(http.MethodPost)

	router.HandleFunc("/item-type", controllers.ListItemType(deps)).Methods(http.MethodGet)
	router.HandleFunc("/item-type", controllers.CreateItemType(deps)).Methods(http.MethodPost)
	router.HandleFunc("/item-types/{name}", controllers.GetItemType(deps)).Methods(http.MethodGet)
	router.HandleFunc("/item-types/{name}", controllers.UpdateItemType(deps)).Methods(http.MethodPost)
	router.HandleFunc("/item-types/{name}", controllers.DeleteItemType(deps)).Methods(http.MethodDelete)

	router.HandleFunc("/event", controllers.CreateEventHandler()).Methods(http.MethodPost)
	router.HandleFunc("/event", controllers.ListEventHandler(deps)).Methods(http.MethodGet)
	router.HandleFunc("/events/{eventId}", controllers.DeleteEventHandler()).Methods(http.MethodDelete)
	router.HandleFunc("/events/{eventId}", controllers.GetEventHandler()).Methods(http.MethodGet)
	router.HandleFunc("/events/{eventId}", controllers.UpdateEventHandler()).Methods(http.MethodPost)

	router.HandleFunc("/user", controllers.CreateUserHandler(deps)).Methods(http.MethodPost)
	router.HandleFunc("/users/{userId}", controllers.GetUserHandler(deps)).Methods(http.MethodGet)

	//router.HandleFunc("/event/{eventId}")
	c := cors.New(cors.Options{
		AllowedOrigins: []string{
			"http://localhost:14045",
			"http://localhost:5173",
		},
		AllowCredentials: true,
		Debug:            true,
	})
	handler := c.Handler(router)
	logger.Infof("Starting server at port %s", config.GetPort())
	server := &http.Server{Handler: handler, Addr: ":" + config.GetPort()}

	go func() {
		err := server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			panic(err)
		}
	}()
	<-stop

	logger.Info("Shutting down server...")
	server.Shutdown(context.Background())

}
