package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zawadimario/mysql-crud-api/pkg/database"
	"github.com/zawadimario/mysql-crud-api/pkg/models"
)

var db *sql.DB

func GetAlbumByArtist(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	artists := vars["artist"]
	albums, err := albumsByArtist(artists, database.Conn)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error retrieving albums: %v", err), http.StatusInternalServerError)
		return
	}
	response := struct {
		Albums []models.Album `json:"albums"`
	}{
		Albums: albums,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error encoding JSON response: %v", err), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

// albumsByArtist queries for albums that have the specified artist name.
func albumsByArtist(name string, db *sql.DB) ([]models.Album, error) {
	// An albums slice to hold data from returned rows.
	var albums []models.Album

	rows, err := db.Query("SELECT * FROM recordings.album WHERE Artist = ?", name)
	if err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var alb models.Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price, &alb.LastUpdate); err != nil {
			return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
		}
		albums = append(albums, alb)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	return albums, nil
}
