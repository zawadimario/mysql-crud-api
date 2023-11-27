package main

import (
	"example/data-access/pkg/database"
	"example/data-access/pkg/handlers"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	db := database.Init()

	defer database.CloseDB(db)

	router := mux.NewRouter()

	// Define your routes
	router.HandleFunc("/albums", handlers.GetAlbums).Methods("GET")
	router.HandleFunc("/albums/{id}", handlers.GetAlbumById).Methods("GET")
	router.HandleFunc("/albums/artist/{artist}", handlers.GetAlbumByArtist).Methods("GET")
	router.HandleFunc("/albums", handlers.CreateAlbum).Methods("POST")
	router.HandleFunc("/albums/{id}", handlers.UpdateAlbum).Methods("PUT")
	router.HandleFunc("/albums/{id}", handlers.DeleteAlbum).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", router))
}
