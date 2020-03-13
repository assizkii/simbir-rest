package rest

import (
	"github.com/assizkii/simbir-rest/internal/adapters/storages"
	"github.com/assizkii/simbir-rest/pkg/rest/handlers"
	"github.com/assizkii/simbir-rest/pkg/rest/utils"
	"log"
	"net/http"
	"os"
)

func RunServer() {

	conf := utils.InitConfig()
	//dsn := "host=0.0.0.0 user=simbir password=simbir dbname=simbir_db  sslmode=disable"
	storage := storages.NewPgStorage(conf.Database)
	handler := handlers.Init(&storage)
	mux := http.NewServeMux()

	port := os.Getenv("PORT")
	if port == "" {
		port = "3030" //localhost
	}

	mux.HandleFunc("/reg", handler.Registration)
	mux.HandleFunc("/auth", handler.Auth)
	mux.HandleFunc("/logout", handler.Logout)
	mux.HandleFunc("/getNumber", handler.GetRandNumber)

	log.Printf("rest is listening at %s", ":5000")

	err := http.ListenAndServe(conf.Host, mux)

	if err != nil {
		log.Fatal("ListenAndServe: ", conf.Host)
	}
}
