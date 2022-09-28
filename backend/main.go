package main

import (
	"backend/pkg/models"
	"math/rand"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	// "github.com/google/uuid"
	_ "github.com/joho/godotenv/autoload"
)

var (
	HOST      = os.Getenv("HOST")
	DB_URL    = os.Getenv("POSTGRE_URL")
	DbUser    = map[string]models.User{}
	DbSession = map[string]string{}
)

func RNG() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(100000)
}

func main() {
	router := gin.Default()
	// db := database.Connect(DB_URL)
	// log.Println(db.Config)

	//ALllows cross site stuff for local development, REMOVE BEFORE PROD
	corsOpts := cors.DefaultConfig()
	corsOpts.AllowAllOrigins = true
	router.Use(cors.New(corsOpts))

	router.GET("/rng", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"NUMBER": RNG(),
		})
	})

	router.Run(HOST)
}
