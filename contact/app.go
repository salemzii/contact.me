package contact

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
)

var contactRespository *SQLiteRepository

func init() {

	db, err := sql.Open("sqlite3", "contact.db")
	if err != nil {
		log.Fatal(err)
	}
	contactRespository = NewSqliteRepository(db)
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

/*
{
  "name":"salemi",
  "email":"salemi@gmail.com",
  "phone_number":"80959023",
  "home_address":"kano"
}
*/

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

func LookupContact(c *gin.Context) {
	name := c.Param("name")

	foundContact, err := contactRespository.GetByName(name)
	if err != nil {
		c.JSON(404, gin.H{
			"error": err,
		})
	}

	c.JSON(200, gin.H{
		"id":           foundContact.ID,
		"name":         foundContact.Name,
		"email":        foundContact.Email,
		"phone_number": foundContact.PhoneNumber,
		"address":      foundContact.HomeAddress,
	})

}
