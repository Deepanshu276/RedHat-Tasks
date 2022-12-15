package router

import (
	"github.com/Deepanshu276/mongoapi/controller"
	//Gorrila mux is used to establish different Routes
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	//Creating the mux router
	router := mux.NewRouter()
	//Fetching all the movies from the Database and we do "get" request for this
	router.HandleFunc("/api/movies", controller.GetMyAllMovies).Methods("GET")
	//adding the movie into the Database and we do "post" request for this
	router.HandleFunc("/api/movie", controller.CreateMovie).Methods("POST")
	//Updating the Status of weather movie is watched or not and we use "put" request for it
	router.HandleFunc("/api/movie/{id}", controller.MarkAsWatched).Methods("PUT")
	//Delete one movie and we use "Delete" request for it
	router.HandleFunc("/api/movie/{id}", controller.DeleteAMovie).Methods("DELETE")
	//Delete All the movie present in the Database and we use "Delete" request for this
	router.HandleFunc("/api/deleteallmovie", controller.DeleteAllMovies).Methods("DELETE")
	//This will return reference of mux router used in main.go file
	return router
}
