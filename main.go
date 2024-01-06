package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// /////////////////////////////////////////db
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

///////////////////////////////////////////

func getOwner(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, CCs)

}

func errormes() {
	fmt.Println("User not found!") //used in func below
}

func getOwnerById(c *gin.Context) {
	ID := c.Param("ID")
	for _, a := range CCs {
		if a.ID == ID {
			c.IndentedJSON(http.StatusOK, a)
		} else {
			c.IndentedJSON(http.StatusBadRequest, errormes)
		}
	}

} //code doesnt work when this func is in it?????????

func createOwner(c *gin.Context) {
	var newOwner CreditCard

	if err := c.BindJSON(&newOwner); err != nil {
		return
	}

	CCs = append(CCs, newOwner)
	c.IndentedJSON(http.StatusCreated, newOwner)
}

func deleteOwner(c *gin.Context) {
	ID := c.Param("ID")
	client := &http.Client{}
	req, err := http.NewRequest("DELETE", "http://localhost:8080/CCs/"+ID, nil)
	if err != nil {
		panic(err)
	}

	// send the request
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// set the response status code
	c.Status(resp.StatusCode)
}

func main() {
	r := gin.Default()
	public := r.Group("/api")
	public.GET("/CCs", getOwner)         //works
	public.GET("/CCs/:ID", getOwnerById) //works
	public.POST("/CCs", createOwner)     //works
	//r.PUT("/CCs/:ID", updateOwnerData) //nw
	public.DELETE("/CCs/:ID", deleteOwner)
	r.Run("localhost:8080")
}
