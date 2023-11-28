package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/zawadimario/mysql-crud-api/pkg/database"
	"github.com/zawadimario/mysql-crud-api/pkg/handlers"
)

func main() {
	db := database.Init()
	defer database.CloseDB(db)

	router := mux.NewRouter()

	// Define your routes
	router.HandleFunc("/albums", handlers.GetAlbums).Methods(http.MethodGet)
	router.HandleFunc("/albums/{id}", handlers.GetAlbumById).Methods(http.MethodGet)
	router.HandleFunc("/albums/artist/{artist}", handlers.GetAlbumByArtist).Methods(http.MethodGet)
	router.HandleFunc("/albums", handlers.CreateAlbum).Methods(http.MethodPost)
	router.HandleFunc("/albums/{id}", handlers.UpdateAlbum).Methods(http.MethodPost)
	router.HandleFunc("/albums/{id}", handlers.DeleteAlbum).Methods(http.MethodDelete)

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", router))
}
