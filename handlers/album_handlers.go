package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Artemych91/web-service-gin/database"
	"github.com/Artemych91/web-service-gin/models"
	"github.com/gin-gonic/gin"
)

const (
	StatusInternalServerError = http.StatusInternalServerError
	StatusBadRequest          = http.StatusBadRequest
	StatusNotFound            = http.StatusNotFound
	StatusCreated             = http.StatusCreated
	StatusOK                  = http.StatusOK
)

// getAlbums responds with the list of all albums as JSON.
func GetAlbums(c *gin.Context) {
	db, err := database.GetDB()
	if err != nil {
		HandleError(c, StatusInternalServerError, err)
		return
	}

	rows, err := db.Query("SELECT * FROM albums")
	if err != nil {
		HandleError(c, StatusInternalServerError, err)
		return
	}

	var albums []models.Album
	for rows.Next() {
		var a models.Album
		err := rows.Scan(&a.ID, &a.Title, &a.Artist, &a.Price)
		if err != nil {
			HandleError(c, StatusInternalServerError, err)
			return
		}
		albums = append(albums, a)
	}

	c.IndentedJSON(StatusOK, albums)
}

func GetAlbum(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		HandleError(c, StatusBadRequest, errors.New("Invalid album ID"))
		return
	}

	db, err := database.GetDB()
	if err != nil {
		HandleError(c, StatusInternalServerError, err)
		return
	}

	row := db.QueryRow("SELECT * FROM albums WHERE id = ?", id)
	var a models.Album
	err = row.Scan(&a.ID, &a.Title, &a.Artist, &a.Price)
	if err != nil {
		HandleError(c, StatusNotFound, errors.New("Album not found"))
		return
	}

	c.IndentedJSON(StatusOK, a)
}

func PostAlbum(c *gin.Context) {
	var newAlbum models.Album
	err := c.BindJSON(&newAlbum)
	if err != nil {
		HandleError(c, StatusInternalServerError, err)
		return
	}

	db, err := database.GetDB()
	if err != nil {
		HandleError(c, StatusInternalServerError, err)
		return
	}

	result, err := db.Exec("INSERT INTO albums (title, artist, price) VALUES (?, ?, ?)", newAlbum.Title, newAlbum.Artist, newAlbum.Price)
	if err != nil {
		HandleError(c, StatusInternalServerError, err)
		return
	}

	id, _ := result.LastInsertId()
	newAlbum.ID = strconv.FormatInt(id, 10)

	c.IndentedJSON(http.StatusCreated, newAlbum)
}
