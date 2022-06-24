package contact

import (
	"database/sql"
	//"html/template"
	"log"
	"net/http"
	"fmt"

	"github.com/gin-gonic/gin"
	_"github.com/lib/pq"
)

var contactRespository *PostgresRepository

func init() {

	uri := fmt.Sprintf("%s://%s:%s@%s/%s?sslmode=disable", "postgresql", "postgres", "auth1234", "localhost:5432", "contact")

	db, err := sql.Open("postgres", uri)
	if err != nil {
		log.Fatal(err)
	}

	contactRespository = NewPostgresRepository(db)
	if err := contactRespository.Migrate(); err != nil {
		log.Fatal(err)
	}

}

type ContactInfo struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	HomeAddress string `json:"home_address"`
}

type AddContactResp struct {
	Response string `json:"response"`
}


func Index(c *gin.Context){

	contacts, err := contactRespository.All()
	if err != nil {
		log.Println("error")
		log.Println(err)
	}
		c.HTML(http.StatusOK, "index.html", gin.H{
				"title": "Main website",
				"contacts": contacts,
		})
}

// Post request service to handle incoming contact data
func AddContact(c *gin.Context) {

	var contact ContactInfo

	c.BindJSON(&contact)
	savedContact, err := contactRespository.Create(contact)

	if err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
	}

	fmt.Println(savedContact.Email)
	c.JSON(201, gin.H{
		"id":           savedContact.ID,
		"name":         savedContact.Name,
		"email":        savedContact.Email,
		"phone_number": savedContact.PhoneNumber,
		"address":      savedContact.HomeAddress,
	})

}

type LookupContactByName struct {
	Name string `json:"name"`
}

func AllContacts(c *gin.Context) {

	contacts, err := contactRespository.All()
	if err != nil {
		c.JSON(404, gin.H{
			"error": err,
		})
	}

	c.JSON(200, gin.H{
		"data": contacts,
	})

}

func LookupContact(c *gin.Context) {
	name := c.Param("name")

	foundContact, err := contactRespository.GetByName(name)
	if err != nil {
		c.JSON(404, gin.H{
			"error": err,
		})
	}

	c.HTML(200,  "contact.html", gin.H{
		"id":           foundContact.ID,
		"name":         foundContact.Name,
		"email":        foundContact.Email,
		"phone_number": foundContact.PhoneNumber,
		"address":      foundContact.HomeAddress,
	})

}



func DeleteContact(c *gin.Context){
	name := c.Param("name")
	cont, err := contactRespository.GetByName(name)
	if err != nil {
		c.JSON(404, gin.H{
			"error": err,
		})
	}

	err = contactRespository.Delete(cont.ID); if err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
	}

	c.JSON(200, gin.H{
		"success": "Contact deleted successfully",
	})
}