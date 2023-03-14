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
	Id       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.Id = strconv.Itoa(rand.Intn(1000))
	movies = append(movies, movie)

	json.NewEncoder(w).Encode(movie)
}

//! In this case, we delete the item from data and then add it 
//! again once updated because we are not working with a database

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	// match the id of the movie wanted
	for i, item := range movies {
		if item.Id == params["id"] {
			movies = append(movies[:i], movies[i + 1:]...)	// delete the movie
		}
	}

	// new var movie with the data updated
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movies = append(movies, movie)	// add the updated movie to the data

	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, item := range movies {
		if item.Id == params["id"] {
			movies = append(movies[:i], movies[i + 1:]...)	//? deleting movie from map
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func main() {
	r := mux.NewRouter()

	//* adding random data to use in the api
	movies = append(movies, Movie{
		Id: "1",
		Isbn: "438227",
		Title: "The Batman",
		Director: &Director{
			Firstname: "Stephen",
			Lastname: "Hawking",
		},
	})
	movies = append(movies, Movie{
		Id: "2",
		Isbn: "253954",
		Title: "Scary Movie",
		Director: &Director{
			Firstname: "Jason",
			Lastname: "Spielberg",
		},
	})


	//? GET METHODS
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")

	//? POST METHODS
	r.HandleFunc("/movies", createMovie).Methods("POST")

	//? PUT METHODS
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")

	//? DELETE METHODS
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Starting server at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
