package router

import (
	"mongodbapi/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/movies",controller.GetAllMovies).Methods("GET")
	router.HandleFunc("/api/movie",controller.CreateAMovie).Methods("POST")
	router.HandleFunc("/api/movie/{movieId}",controller.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/movie/{movieId}",controller.DeleteOneMovie).Methods("DELETE")
	router.HandleFunc("/api/movies",controller.DeleteAllMovies).Methods("DELETE")

	return router
}