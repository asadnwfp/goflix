package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/asadnwfp/movie-directory/logger"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type Movie struct {
	Id       string    `json:"id"`
	Title    string    `json:"title"`
	Isbn     string    `json:"isbn"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var movies []Movie

var zapLogger *zap.Logger

func main() {
	zapLogger = logger.GetLogger()

	r := mux.NewRouter()

	movies = append(movies, Movie{Id: "1", Title: "hello", Isbn: "1111", Director: &Director{FirstName: "Saad", LastName: "Ahmed"}})
	movies = append(movies, Movie{Id: "2", Title: "world", Isbn: "2222", Director: &Director{FirstName: "Durrani", LastName: "sajid"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movie/{id}", getMovie).Methods("GET")

	fmt.Print("Listen and Serve on Port :8000\n")
	zapLogger.Info("Server Listening and Serving",
		zap.Int("port", 8000),
	)
	log.Fatal(http.ListenAndServe("0.0.0.0:8000", r))
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	json.NewEncoder(w).Encode(movies)

	jsonData, _ := json.Marshal(movies)
	zapLogger.Info("List of Movies",
		zap.String("json", string(jsonData)))
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	for _, movie := range movies {
		if movie.Id == params["id"] {
			json.NewEncoder(w).Encode(movie)
		}
	}

}
