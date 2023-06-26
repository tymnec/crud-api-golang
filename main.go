package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

// Get Movies Function
func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

// Main Function
func main() {
	r := mux.NewRouter()

	// Appending some movies
	// First Movie
	movies = append(movies, Movie{ID: "1", Isbn: "82476", Title: "Casual Movie", Director: &Director{Firstname: "Nome", Lastname: "Due"}})

	// Second Movie
	movies = append(movies, Movie{ID: "2", Isbn: "29397", Title: "Dada Lakhmi", Director: &Director{Firstname: "Jem", Lastname: "Mia"}})

	// Create 5 Routes for each Function
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovies).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovies).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovies).Methods("DELETE")

	fmt.Printf("Starting the server at 8080\n")
	log.Fatal(http.ListenAndServe(":8080", r))
}
