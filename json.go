package main

import (
	"encoding/json"
	"log"
	"net/http"
)
func respondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Printf("Responding with 5XX error: %s", msg)
	}
	type errorResponse struct {
		Error string `json:"error"`
	}
	//rhis function write the json to the header
	respondWithJSON(w, code, errorResponse{
		Error: msg,
	})
}
func respondWithJSON(w http.ResponseWriter , code int, payload interface{}){
	data, err := json.Marshal(payload)
	if err!= nil {
		log.Fatal("Error marshalling the payload to json")
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)

}