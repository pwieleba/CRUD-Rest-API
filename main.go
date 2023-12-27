package main

import (
	"net/http"
	//"github.com/CRUD-Rest-API/api"
	"github.com/gin-gonic/gin"
)

func getOwner(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, CCs)

}

//func errormes() {
//	fmt.Println("User not found!") //used in func below
//}

//func getOwnerById(c *gin.Context) {
//ID := c.Param("ID")
//for _, a := range CCs {
//if a.ID == ID {
//c.IndentedJSON(http.StatusOK, a)
//	} //else if c.IndentedJSON(http.StatusBadRequest, errormes);
// }

//} //code doesnt work when this func is in it?????????

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
	r.GET("/CCs", getOwner) //works
	//r.GET("/CCs/:ID", getOwnerById) //do not works
	r.POST("/CCs", createOwner) //works
	// r.PUT("/CCs/:ID", updateOwnerData) //nw
	r.DELETE("/CCs/:ID", deleteOwner) //?
	r.Run("localhost:8080")
}
