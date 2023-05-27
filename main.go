package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"direcotor"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1",Isbn: "784935",Title: "first movie",Director: &Director{Lastname: "test",Firstname: "name"}})
	movies = append(movies, Movie{ID: "2",Isbn: "7834535",Title: "second movie",Director: &Director{Lastname: "test2",Firstname: "name2"}})
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/${id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/${id}", updateMovie).Methods("POST")
	r.HandleFunc("/movies/${id}", deleteMovie).Methods("DELETE")

	fmt.Print("starting server at port 8000 \n")
	err := http.ListenAndServe(":8000", r)
	if err != nil {
		log.Fatal(err)
	}
}
