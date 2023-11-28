package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/gommon/log"
	"github.com/zawadimario/mysql-crud-api/pkg/database"
	"github.com/zawadimario/mysql-crud-api/pkg/models"
)

func GetAlbums(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	albums, err := getAllAlbums(database.Conn)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error retrieving albums: %v", err), http.StatusInternalServerError)
		return
	}

	response := struct {
		Albums []models.Album
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
func getAllAlbums(db *sql.DB) ([]models.Album, error) {
	var albums []models.Album

	rows, err := db.Query("SELECT * FROM recordings.album ORDER BY ID")
	if err != nil {
		return nil, fmt.Errorf("getAllAlbums: %v", err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Errorf("closing connection :%v", err)
		}
	}(rows)

	for rows.Next() {
		var alb models.Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price, &alb.LastUpdate); err != nil {
			return nil, fmt.Errorf("getAllAlbums: %v", err)
		}
		albums = append(albums, alb)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("getAllAlbums: %v", err)
	}
	return albums, nil
}
