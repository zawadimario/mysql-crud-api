package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/zawadimario/mysql-crud-api/pkg/database"
)

func DeleteAlbum(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid album ID", http.StatusBadRequest)
		return
	}
	err = deleteAlbumByID(int64(id), database.Conn)
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
