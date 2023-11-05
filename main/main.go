package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	controll.getOwner()
	controll.getOwnerById()
	controll.createOwner()
	controll.updateOwnerData()
	controll.deleteOwner()

	r := gin.Default()
	r.GET("/CCs", getOwner)
	r.GET("/CCs/:id", getOwnerById)
	r.POST("/CCs", createOwner)
	r.PUT("/CCs/:id", updateOwnerData)
	r.DELETE("/CCs/:id", deleteOwner)
	r.Run("localhost:8080")
}
