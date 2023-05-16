package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Struct tags such as json:"artist" specify what a field’s name should be when the struct’s contents are serialized into JSON.
// Without them, the JSON would use the struct’s capitalized field names – a style not as common in JSON.

// album represents data about a record album.
type album struct {
	Id     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{Id: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{Id: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{Id: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.GET("/albums/:id", getAlbumById)

	_ = router.Run("localhost:8080")
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum album

	if err := c.ShouldBindJSON(&newAlbum); err != nil {
		// Call BindJSON to bind the received JSON to newAlbum.
		_ = c.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumById locates the album whose id value matches the id parameter sent by the client,
// then returns that album as a response
func getAlbumById(c *gin.Context) {
	id := c.Param("id")

	// Loop over the slice of albums, looking for
	// an album whose id value matches the parameter.
	for _, element := range albums {
		if element.Id == id {
			c.IndentedJSON(http.StatusOK, element)
			return
		}
	}
	c.IndentedJSON(http.StatusNoContent, nil)
}
