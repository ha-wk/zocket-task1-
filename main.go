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
	Title    string    `json:"title"` //MOVIES ATTRIBUTES
	Director *Director `json:"director"`
}
type Director struct {
	FirstName string `json:"firstname"` //MOVIES ATTRIBUTES
	LastName  string `json:"lastname"`
}

var movies []Movie //DATA STRUCTURE TO HOLD MOVIES AND ITS DETAILS

func getMovies(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json") //TO FETCH ALL MOVIES(READ)
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //TO DELETE A MOVIE(DELETE)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {

		if item.ID == params["id"] { //TO GET A SINGLE MOVIE(GET)
			json.NewEncoder(w).Encode(item)
			return
		}
	}

}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie) //TO CREATE A NEW MOVIE(CREATE)
	movie.ID = strconv.Itoa(rand.Intn(999999))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)

}

func updateMovie(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...) //TO UPDATE ANY MOVIE ATTRIBUTES(UPDATE)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

func main() {
	r := mux.NewRouter()

	//SAMPLE DUMMIE MOVIE DATA....WE CAN ALSO USE ANY DATABASE FOR SAME!
	movies = append(movies, Movie{ID: "1", Isbn: "222222", Title: "Movie 1", Director: &Director{FirstName: "John", LastName: "Wick"}})
	movies = append(movies, Movie{ID: "2", Isbn: "444444", Title: "Movie 2", Director: &Director{FirstName: "Larry", LastName: "Brin"}})

	//ALL REQUIRED ENDPOINTS
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	//ASSIGNING PORT WHERE OUR SERVER WILL RUN....
	fmt.Printf("staring server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))

}
