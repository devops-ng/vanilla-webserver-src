package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/scottwinkler/vanilla-webserver-src/server/action/pets"
	"github.com/scottwinkler/vanilla-webserver-src/server/model/pet"
)

var db *gorm.DB

func init() {
	initializeRDSConn()
	validateRDS()
}

func initializeRDSConn() {
	//read from configuration file
	dat, _ := ioutil.ReadFile("/etc/server.conf")
	fmt.Printf("%s", string(dat))
	m := make(map[string]string)
	json.Unmarshal(dat, &m)
	user := m["pg_user"]         //admin
	password := m["pg_password"] //PJwuu-MbCsEXdU__
	netloc := m["pg_netloc"]     //my-cool-project-db-instance.cozpurlif6yt.us-west-2.rds.amazonaws.com:3306
	database := m["pg_database"] //pets

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", user, password, netloc, database)
	var err error
	db, err = gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("%s", err)
	}
}

func validateRDS() {
	//If the pets table does not already exist, create it
	if !db.HasTable("pets") {
		db.CreateTable(&pet.Pet{})
	}
}

func main() {
	r := gin.Default()
	r.Use(static.Serve("/", static.LocalFile("/run/public", false)))
	r.POST("/pets", createPetHandler)
	r.GET("/pets/:id", getPetHandler)
	r.GET("/pets", listPetsHandler)
	r.Run()
}

func createPetHandler(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	var req pets.CreatePetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := pets.CreatePet(db, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
	return
}

func listPetsHandler(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	limit := 10
	if c.Query("limit") != "" {
		newLimit, err := strconv.Atoi(c.Query("limit"))
		if err != nil {
			limit = 10
		} else {
			limit = newLimit
		}
	}
	if limit > 50 {
		limit = 50
	}
	req := pets.ListPetsRequest{Limit: uint(limit)}
	res, _ := pets.ListPets(db, &req)
	c.JSON(http.StatusOK, res)
}

func getPetHandler(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	id := c.Param("id")
	req := pets.GetPetRequest{ID: id}
	res, _ := pets.GetPet(db, &req)
	if res.Pet == nil {
		c.JSON(http.StatusNotFound, res)
		return
	}
	c.JSON(http.StatusOK, res)
}
