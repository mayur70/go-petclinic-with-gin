package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.tmpl")
	router.GET("/", welcome)
	router.GET("/oups", triggerException)
	router.GET("/vets", showResourcesVetList)
	router.GET("/vets.html", showVetList)

	router.Run("localhost:8080")
}

type Vet struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"lastName"`
}

var vets = []Vet{
	{ID: 1, FirstName: "Mayur", LastName: "Patil"},
}

func welcome(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "welcome"})
}

func triggerException(c *gin.Context) {
	//TODO: throw expection??
}

func showVetList(c *gin.Context) {
	c.HTML(http.StatusOK, "vets.tmpl", gin.H{
		"title": "vets page is HERE",
	})
}

func showResourcesVetList(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, vets)
}

func findVetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
		return
	}

	for _, v := range vets {
		if v.ID == id {
			c.IndentedJSON(http.StatusOK, v)
			return
		}
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "vet not found"})
}
