package controll

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

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
	uID := c.Param("ID")
	uBalance := c.Param("Balance")
	uOwner := c.Param("Owner")

}

func deleteOwner(c *gin.Context) {

}
