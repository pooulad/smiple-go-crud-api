package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

func getMovies(res http.ResponseWriter, req *http.Request) {
	fmt.Print("movies called")
	res.Header().Set("Content-type", "application/json")
	json.NewEncoder(res).Encode(movies)
}

func deleteMovie(res http.ResponseWriter, req *http.Request) {
	fmt.Print("movies delete called")
	res.Header().Set("Content-type", "application/json")
	params := mux.Vars(req)
	for index, value := range movies {
		if value.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(res).Encode(movies)
	json.NewEncoder(res).Encode(movies)
}

func getMovie(res http.ResponseWriter, req *http.Request) {
	fmt.Print("movies id called")
	res.Header().Set("Content-type", "application/json")
	params := mux.Vars(req)
	for _, value := range movies {
		if value.ID == params["id"] {
			json.NewEncoder(res).Encode(value)
			return
		}
	}
}

func createMovie(res http.ResponseWriter, req *http.Request) {
	fmt.Print("create movies called")
	res.Header().Set("Content-type", "application/json")
	var movie Movie
	_ = json.NewDecoder(req.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(1000000000))
	movies = append(movies, movie)
	json.NewEncoder(res).Encode(movie)
}

func updateMovie(res http.ResponseWriter, req *http.Request) {
	fmt.Print("update movies called")
	res.Header().Set("Content-type", "application/json")
	params := mux.Vars(req)
	for index, value := range movies {
		if value.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(req.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(res).Encode(movie)
			return
		}
	}
}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "784935", Title: "first movie", Director: &Director{Lastname: "test", Firstname: "name"}})
	movies = append(movies, Movie{ID: "2", Isbn: "7834535", Title: "second movie", Director: &Director{Lastname: "test2", Firstname: "name2"}})
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Print("starting server at port 8000 \n")
	err := http.ListenAndServe(":8000", r)
	if err != nil {
		log.Fatal(err)
	}
}
