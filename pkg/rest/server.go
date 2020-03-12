package rest

import (
	"log"
	"net/http"
	"os"
	"github.com/assizkii/simbir-rest/pkg/rest/utils"
)

func RunServer()  {

	conf := utils.InitConfig()

	mux := http.NewServeMux()

	//router.NotFoundHandler = app.NotFoundHandler

	port := os.Getenv("PORT")
	if port == "" {
		port = "3030" //localhost
	}

	log.Printf("rest is listening at %s", conf.Host)

	err := http.ListenAndServe(conf.Host, mux)

	if err != nil {
		log.Fatal("ListenAndServe: ", conf.Host)
	}
}