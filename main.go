package main

import (
	// "fmt"
	"log"
	"net/http"
	"os"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main(){
	godotenv.Load()

	port := os.Getenv("PORT")
	
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	server := &http.Server{
		Handler: router,
		Addr : ":" + port,
	
	}
	v1Router:= chi.NewRouter()
	v1Router.Get("/healthz",handlerReadiness)
	v1Router.Get("/err", handlerErr)
	router.Mount("/v1",v1Router)

	log.Printf("------- Server Starting ------- at Port: %v", port)
	// stops here and starts handling http requests
	err := server.ListenAndServe()
	if err!= nil{
		log.Fatal(err)
	}
}