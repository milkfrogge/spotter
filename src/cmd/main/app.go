package main

import (
	"SpotterBackend/src/internal/config"
	"SpotterBackend/src/pkg/logging"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
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
	log.Printf("App appConfig: %v\n", appConfig)
	log.Infof("Statring application")
	router := httprouter.New()
	router.GET("/", Index)
	log.Fatal(http.ListenAndServe(appConfig.ConnectionConfig.BindAddress, router))
}
