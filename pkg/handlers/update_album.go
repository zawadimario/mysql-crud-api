package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/zawadimario/mysql-crud-api/pkg/models"
)

func UpdateAlbum(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid album ID", http.StatusBadRequest)
		return
	}
	var updatedAlbum models.Album
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&updatedAlbum); err != nil {
		http.Error(w, fmt.Sprintf("Invalid request payload: %v", err), http.StatusBadRequest)
		return
	}
	err = updateAlbumById(id, updatedAlbum, db)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error updating album: %v", err), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func updateAlbumById(id int64, updatedAlbum models.Album, db *sql.DB) error {
	_, err := db.Exec("UPDATE recordings.album SET Title = ?, Artist = ?, Price = ? WHERE ID = ?", updatedAlbum.Title, updatedAlbum.Artist, updatedAlbum.Price, updatedAlbum.ID)
	if err != nil {
		return fmt.Errorf("updateAlbumByID %d: %v", id, err)
	}
	return nil
}
