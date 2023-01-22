package main

import (
	"SpotterBackend/src/internal/config"
	"SpotterBackend/src/internal/user/db"
	"SpotterBackend/src/pkg/client"
	"SpotterBackend/src/pkg/logging"
	"context"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log := logging.GetLogger()
	_, err := fmt.Fprint(w, "Welcome %s!\n", r.Host)
	if err != nil {
		log.Fatal("")
	}
}

func main() {
	log := logging.GetLogger()
	appConfig := config.GetConfig()
	pgClient, err := client.NewClient(context.Background(), appConfig.DbConfig)
	if err != nil {
		log.Fatalf("error while init db: %s", err)
	}
	userStorage := db.NewStorage(pgClient, log)
	userStorage.FindOne(context.Background(), 0)
	log.Printf("App appConfig: %v\n", appConfig)
	log.Infof("Statring application")
	router := httprouter.New()
	router.GET("/", Index)
	log.Fatal(http.ListenAndServe(appConfig.ConnConfig.BindAddress, router))
}
