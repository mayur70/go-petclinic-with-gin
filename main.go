package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/", welcome)
	router.GET("/oups", triggerException)
	router.GET("/vets", showResourcesVetList)

	router.Run("localhost:8080")
}

type Vet struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"lastName"`
}

func welcome(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "welcome"})
}

func triggerException(c *gin.Context) {
	//TODO: throw expection??
}

func showResourcesVetList(c *gin.Context) {
	vets := []Vet{
		{ID: 1, FirstName: "Mayur", LastName: "Patil"},
	}
	c.IndentedJSON(http.StatusOK, vets)
}
