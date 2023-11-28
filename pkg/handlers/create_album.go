package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/zawadimario/mysql-crud-api/pkg/database"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/zawadimario/mysql-crud-api/pkg/models"
)

func CreateAlbum(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newAlbum models.Album

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&newAlbum); err != nil {
		http.Error(w, fmt.Sprintf("Invalid request payload: %v", err), http.StatusBadRequest)
		return
	}
	id, err := addAlbum(newAlbum, database.Conn)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating album: %v", err), http.StatusInternalServerError)
		return
	}
	response := struct {
		ID int64 `json:"id"`
	}{
		ID: id,
	}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error encoding JSON response: %v", err), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)

}
func addAlbum(alb models.Album, db *sql.DB) (int64, error) {
	result, err := db.Exec("INSERT INTO recordings.album (Title, Artist, Price) VALUES (?, ? , ?)", alb.Title, alb.Artist, alb.Price)
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	return id, nil
}
