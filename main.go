package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreditCard struct {
	ID      string `json:"id"`
	Balance int    `json:"balance"`
	Owner   *Owner `json:"owner"`
}

type Owner struct {
	firstname string
	lastname  string
	age       int
	address   string
}

var CCs = []CreditCard{
	{ID: "2446", Balance: 1345141, Owner: &Owner{firstname: "ex1", lastname: "ex1.2", age: 45, address: "NY"}},
	{ID: "96", Balance: 624262, Owner: &Owner{firstname: "ex2", lastname: "ex2.2", age: 69, address: "FL"}},
	{ID: "245", Balance: 4626262, Owner: &Owner{firstname: "ex2", lastname: "ex3.2", age: 33, address: "WA"}},
}

func getOwner(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, CCs)

}

func getOwnerById(c *gin.Context) {
	ID := c.Param("ID")
	for _, a := range CCs {
		if a.ID == ID {
			c.IndentedJSON(http.StatusOK, a)
		}
	}
}

func createOwner(c *gin.Context) {
	var newOwner CreditCard

	if err := c.BindJSON(&newOwner); err != nil {
		return
	}

	CCs = append(CCs, newOwner)
	c.IndentedJSON(http.StatusCreated, newOwner)
}

func updateOwnerData(c *gin.Context) {

}

func deleteOwner() {

}

func main() {
	r := gin.Default()
	r.GET("/CCs", getOwner)
	r.GET("/CCs/:id", getOwnerById)
	r.POST("/CCs", createOwner)
	r.PUT("/CCs/:id", updateOwnerData)
	r.DELETE("/CCs/:id", deleteOwner)
	r.Run("localhost:8080")
}
