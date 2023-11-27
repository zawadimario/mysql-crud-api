package handlers

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func DeleteAlbum(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid album ID", http.StatusBadRequest)
		return
	}
	err = deleteAlbumByID(id, db)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error deleting album: %v", err), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)

}
func deleteAlbumByID(id int64, db *sql.DB) error {
	_, err := db.Exec("DELETE FROM recordings.album WHERE ID = ?", id)
	if err != nil {
		return fmt.Errorf("deleteAbumByID %d: %v", id, err)
	}
	return nil
}
