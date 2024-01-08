package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// /////////////////////////////////////////db
var DB *gorm.DB

func ConnectDB() {
	err := godotenv.Load(".env")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_DATABASE")
	dbUser := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")

	BURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)

	DB, err = gorm.Open(Dbdriver, DBURL)

	if err != nil {
		fmt.Println("Cannot connect to database ", Dbdriver)
		log.Fatal("connection error:", err)
	} else {
		fmt.Println("We are connected to the database ", Dbdriver)
	}

	DB.AutoMigrate(&User{})

}

// user for token

type User struct {
	gorm.Model
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Password string `gorm:"size:255;not null" json:"password"`
}

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
	{ID: "2446", Balance: 1345141, Owner: &Owner{firstname: "Joe", lastname: "Cotton", age: 45, address: "NY"}},
	{ID: "96", Balance: 624262, Owner: &Owner{firstname: "Donald", lastname: "Trump", age: 69, address: "FL"}},
	{ID: "245", Balance: 4626262, Owner: &Owner{firstname: "Kanye", lastname: "West", age: 33, address: "WA"}},
}

//////////////////// auth

type SingUpInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func SingUp(c *gin.Context) {
	var input SingUpInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "validated"})
}

/////////////////////////// crud

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

}

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
	public.POST("/CCs", createOwner)
	public.POST("/singup", SingUp) //works
	//r.PUT("/CCs/:ID", updateOwnerData) //nw
	public.DELETE("/CCs/:ID", deleteOwner)
	r.Run("localhost:8080")
}
