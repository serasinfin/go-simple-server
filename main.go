package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Movie struct {
	ID       string   `json: "id"`
	Isbn     string   `json: "isbn"`
	Title    string   `json: "title"`
	Director Director `json: "director"`
}

type Director struct {
	Firstname string `json: "firstname"`
	Lastname  string `json: "lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	var found bool
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			found = true
			break
		}
	}
	if found {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Movie deleted successfully"))
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Movie not found"))
	}
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	var found bool
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			w.WriteHeader(http.StatusOK)
			found = true
			return
		}
	}
	if !found {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Movie not found"))
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	var movie Movie
	json.NewDecoder(r.Body).Decode(&movie)
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Movie created successfully"))
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	var movie Movie
	json.NewDecoder(r.Body).Decode(&movie)
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies[index] = movie
			json.NewEncoder(w).Encode(movie)
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Movie updated successfully"))
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Movie not found"))
}

func main() {
	movies = append(movies, Movie{ID: "1", Isbn: "448743", Title: "Golang programming language", Director: Director{Firstname: "John", Lastname: "Smith"}})
	movies = append(movies, Movie{ID: "2", Isbn: "448744", Title: "Golang programming language 2", Director: Director{Firstname: "John", Lastname: "Smith"}})
	movies = append(movies, Movie{ID: "3", Isbn: "448745", Title: "Golang programming language 3", Director: Director{Firstname: "John", Lastname: "Smith"}})

	r := mux.NewRouter()
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Starting server on port 8001")
	log.Fatal(http.ListenAndServe(":8001", r))

}
