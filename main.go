package main

import (
	"github.com/gin-gonic/gin"
)

type CreditCard struct {
	ID      int    `json:"id"`
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
	{ID: 2446, Balance: 1345141, Owner: &Owner{firstname: "ex1", lastname: "ex1.2", age: 45, address: "NY"}},
	{ID: 96, Balance: 624262, Owner: &Owner{firstname: "ex2", lastname: "ex2.2", age: 69, address: "FL"}},
	{ID: 245, Balance: 4626262, Owner: &Owner{firstname: "ex2", lastname: "ex3.2", age: 33, address: "WA"}},
}

func getOwner() {

}

func getOwnerById() {

}

func createOwner() {

}

func updateOwnerData() {

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
