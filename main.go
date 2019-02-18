package tictacgo

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Endpoint struct {
	Path        string      `json:"path"`
	Method      string      `json:"method"`
	Parameters  []Parameter `json:"params,omitempty"`
	Description string      `json:"description,omitempty"`
}

type Parameter struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

var endpoints []Endpoint

func Default(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(endpoints)
}

// our main function
func main() {
	endpoints = append(endpoints, Endpoint{
		Path:        "/",
		Method:      "GET",
		Description: "This endpoint. Lists all of the possible endpoints."})

	/*
	   endpoints = append(endpoints, Endpoint{
	   	Path:   "/game",
	   	Method: "GET",
	   	Parameters: []Parameter{
	   		Parameter{Name: "game-id"}},
	   	Description: "This endpoint. Lists all of the possible endpoints."})
	*/
	router := mux.NewRouter()
	router.HandleFunc("/", Default).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}
