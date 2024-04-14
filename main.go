package main

import (
	"context"
	"errors"
	"github.com/arung-agamani/mitsukeru-go/config"
	"github.com/arung-agamani/mitsukeru-go/controllers"
	db2 "github.com/arung-agamani/mitsukeru-go/db"
	"github.com/arung-agamani/mitsukeru-go/utils/logger"
	"github.com/arung-agamani/mitsukeru-go/utils/validator"
	"github.com/gorilla/mux"
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
	db2.InitDb()
	validator.Init()
	router := mux.NewRouter()
	router.HandleFunc("/", controllers.HelloHandler()).Methods(http.MethodGet)
	router.HandleFunc("/info", controllers.InfoHandler()).Methods(http.MethodGet)
	router.HandleFunc("/item/{itemId}", controllers.GetItemHandler()).Methods(http.MethodGet)
	router.HandleFunc("/event", controllers.CreateEventHandler()).Methods(http.MethodPost)
	router.HandleFunc("/event", controllers.DeleteEventHandler()).Methods(http.MethodDelete)
	router.HandleFunc("/event/{eventId}", controllers.GetEventHandler()).Methods(http.MethodGet)
	router.HandleFunc("/event/{eventId}", controllers.UpdateEventHandler()).Methods(http.MethodPost)
	//router.HandleFunc("/event/{eventId}")

	logger.Infof("Starting server at port %s", config.GetPort())
	server := &http.Server{Handler: router, Addr: ":" + config.GetPort()}

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
