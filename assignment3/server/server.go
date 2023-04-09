package pkg

import (
	"encoding/json"
	"fmt"
	"github.com/isolent/pkg"
	"net/http"
)

var db = database.GetDB()

func getAllMovies(w http.ResponseWriter, r *http.Request) {
	var movies []models.APIMovie
	db.Model(&models.Movie{}).Find(&movies)
	json.NewEncoder(w).Encode(movies)
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	movie := &models.Movie{}
	json.NewDecoder(r.Body).Decode(movie)
	db.Create(movie)
}

func getMovieById(w http.ResponseWriter, r *http.Request) {
	movie := models.APIMovie{}
	id := r.URL.Query().Get("id")
	db.Model(&models.Movie{}).First(&movie, "id = ?", id)
	json.NewEncoder(w).Encode(movie)
}

func RunServer() {
	http.HandleFunc("/movies", getAllMovies)
	http.HandleFunc("/postMovie", createMovie)
	http.HandleFunc("/movies/:id", getMovieById)
	fmt.Println("Server is listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}