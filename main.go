package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/artemych91/web-service-gin/database"
)

const (
	StatusInternalServerError = http.StatusInternalServerError
	StatusBadRequest          = http.StatusBadRequest
	StatusNotFound            = http.StatusNotFound
	StatusCreated             = http.StatusCreated
	StatusOK                  = http.StatusOK
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

func handleError(c *gin.Context, statusCode int, err error) {
	c.IndentedJSON(statusCode, gin.H{"error": err.Error()})
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	db, err := GetDB()
	if err != nil {
		handleError(c, StatusInternalServerError, err)
		return
	}

	rows, err := db.Query("SELECT * FROM albums")
	if err != nil {
		handleError(c, StatusInternalServerError, err)
		return
	}

	var albums []album
	for rows.Next() {
		var a album
		err := rows.Scan(&a.ID, &a.Title, &a.Artist, &a.Price)
		if err != nil {
			handleError(c, StatusInternalServerError, err)
			return
		}
		albums = append(albums, a)
	}

	c.IndentedJSON(StatusOK, albums)
}

func getAlbum(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		handleError(c, StatusBadRequest, errors.New("Invalid album ID"))
		return
	}

	db, err := getDB()
	if err != nil {
		handleError(c, StatusInternalServerError, err)
		return
	}

	row := db.QueryRow("SELECT * FROM albums WHERE id = ?", id)
	var a album
	err = row.Scan(&a.ID, &a.Title, &a.Artist, &a.Price)
	if err != nil {
		handleError(c, StatusNotFound, errors.New("Album not found"))
		return
	}

	c.IndentedJSON(StatusOK, a)
}

func postAlbum(c *gin.Context) {
	var newAlbum album
	err := c.BindJSON(&newAlbum)
	if err != nil {
		handleError(c, StatusInternalServerError, err)
		return
	}

	db, err := getDB()
	if err != nil {
		handleError(c, StatusInternalServerError, err)
		return
	}

	result, err := db.Exec("INSERT INTO albums (title, artist, price) VALUES (?, ?, ?)", newAlbum.Title, newAlbum.Artist, newAlbum.Price)
	if err != nil {
		handleError(c, StatusInternalServerError, err)
		return
	}

	id, _ := result.LastInsertId()
	newAlbum.ID = strconv.FormatInt(id, 10)

	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func main() {
	db, error := initDB()
	if error != nil {
		return
	}
	defer db.Close()

	router := gin.Default()

	// CORS middleware
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	router.Use(cors.New(config))

	// Routers
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbum)
	router.POST("/albums", postAlbum)

	router.Run("localhost:8080")
}
