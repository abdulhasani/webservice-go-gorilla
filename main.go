package main

import (
	"net/http"

	"encoding/json"

	"log"

	"fmt"

	"github.com/gorilla/mux"
)

//Movie Struct
type Movie struct {
	Title  string `json:"title"`
	Rating string `json:"rating"`
	Year   string `json:"year"`
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/movies", handleMovies).Methods("GET")
	http.ListenAndServe(":8080", router)
}

func handleMovies(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	var movies = map[string]*Movie{
		"tt0076759": &Movie{Title: "Star Wars: A New Hope", Rating: "8.7", Year: "1977"},
		"tt0082971": &Movie{Title: "Indiana Jones: Raiders of the Lost Ark", Rating: "8.6", Year: "1981"},
	}
	outgoingJSON, error := json.Marshal(movies)
	if error != nil {
		log.Println(error.Error())
		http.Error(res, error.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprint(res, string(outgoingJSON))
}
