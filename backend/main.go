package main

import (
	"log"
	"math/rand"
	"os"
	"time"

	"backend/pkg/database"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

var (
	HOST   = os.Getenv("HOST")
	DB_URL = os.Getenv("POSTGRE_URL")
)

func RNG() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(100000)
}

func createCookie(c *gin.Context) {
	c.SetCookie("name", "Alex", 10, "/", "", true, true)
	c.JSON(200, gin.H{
		"Cookie": "Created",
	})

}
func readCookie(c *gin.Context) {
	cookie, err := c.Cookie("name")
	if err != nil {
		log.Fatal(err)
		c.Redirect(303, "/")
	}
	c.JSON(200, gin.H{
		"COOKIE:": cookie,
	})
}
func clearcookie(c *gin.Context) {
	c.SetCookie("name", "Alex", -1, "/", "", true, true)
	c.JSON(200, gin.H{
		"COOKIE:": "Cleared",
	})
}

func main() {
	router := gin.Default()
	db := database.Connect(DB_URL)
	log.Println(db.Config)

	//ALllows cross site stuff for local development, REMOVE BEFORE PROD
	corsOpts := cors.DefaultConfig()
	corsOpts.AllowAllOrigins = true
	router.Use(cors.New(corsOpts))
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"NUMBER": RNG(),
		})
	})

	router.GET("/cookie/new", createCookie)
	router.GET("/cookie/read", readCookie)
	router.GET("/cookie/clear", clearcookie)

	router.Run(HOST)
}
