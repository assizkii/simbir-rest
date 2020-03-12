package rest_server

import (
	"log"
	"net/http"
	"os"
	"simbir-rest/pkg/rest-server/utils"
)

func RunServer()  {

	conf := utils.InitConfig()

	mux := http.NewServeMux()

	//router.NotFoundHandler = app.NotFoundHandler

	port := os.Getenv("PORT")
	if port == "" {
		port = "3030" //localhost
	}

	log.Printf("server is listening at %s", conf.Host)

	err := http.ListenAndServe(conf.Host, mux)

	if err != nil {
		log.Fatal("ListenAndServe: ", conf.Host)
	}
}