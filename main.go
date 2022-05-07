package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	"github.com/salemzii/contactInfo/contact"
)

func main() {

	SetUpServer().Run()
}

func SetUpServer() *gin.Engine {
	router := gin.Default()
	router.POST("/add", contact.AddContact)
	router.GET("/lookup/:name", contact.LookupContact)

	return router
}
