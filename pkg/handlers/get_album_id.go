package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/zawadimario/mysql-crud-api/pkg/database"
	"github.com/zawadimario/mysql-crud-api/pkg/models"
)

func GetAlbumById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid album ID", http.StatusBadRequest)
		return
	}
	album, err := albumByID(id, database.Conn)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error retrieving album: %v", err), http.StatusInternalServerError)
		return
	}

	response := struct {
		Album models.Album `json:"album"`
	}{
		Album: album,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error encoding JSON response: %v", err), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)

}

func albumByID(id int64, db *sql.DB) (models.Album, error) {
	var alb models.Album

	row := db.QueryRow("SELECT * FROM recordings.album WHERE id = ?", id)
	if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return alb, fmt.Errorf("albumsById %d: no such album", id)
		}
		return alb, fmt.Errorf("albumById %d: %v", id, err)
	}
	return alb, nil
}
