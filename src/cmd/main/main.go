package main

import (
	_ "SpotterBackend/src/docs"
	"SpotterBackend/src/internal/domain/user/handler"
	"SpotterBackend/src/internal/domain/user/service"
	"SpotterBackend/src/internal/domain/user/storage"
	"SpotterBackend/src/pkg/client"
	"SpotterBackend/src/pkg/config"
	"SpotterBackend/src/pkg/logging"
	"context"
	"fmt"
	"github.com/julienschmidt/httprouter"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log := logging.GetLogger()
	_, err := fmt.Fprint(w, "Welcome %s!\n", r.Host)
	if err != nil {
		log.Fatal("")
	}
}

func swaggerHandler(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
	httpSwagger.WrapHandler(res, req)
}

func main() {
	log := logging.GetLogger()
	appConfig := config.GetConfig()
	router := httprouter.New()
	router.GET("/swagger/:any", swaggerHandler)
	//
	pgClient, err := client.NewClient(context.Background(), appConfig.DbConfig)
	if err != nil {
		log.Fatalf("error while init db: %s", err)
	}
	userStorage := storage.NewStorage(pgClient, log)
	userService := service.NewUserService(userStorage, log)
	userHandler := handler.NewUserHandler(log, userService)

	userHandler.Register(router)
	log.Printf("App appConfig: %v\n", appConfig)
	log.Infof("Statring application")
	router.GET("/", Index)
	log.Fatal(http.ListenAndServe(appConfig.ConnConfig.BindAddress, router))
}
