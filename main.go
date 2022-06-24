package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	"github.com/salemzii/contactInfo/contact"
)

func main() {

	SetUpServer().Run(":8000")
}

func SetUpServer() *gin.Engine {
	router := gin.Default()

	router.LoadHTMLGlob("template/*")
	router.POST("/add", contact.AddContact)
	router.GET("/lookup/:name", contact.LookupContact)
	router.GET("/allcontacts", contact.AllContacts)
	router.DELETE("/delete/:name", contact.DeleteContact)

	router.GET("/index", contact.Index)

	return router
}


