package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/konojunya/twblock-suspicious-account/router"
	"github.com/konojunya/twblock-suspicious-account/service"
)

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	service.SetKeys(os.Getenv("consumerKey"), os.Getenv("consumerSecret"))
}

func main() {
	loadEnv()
	r := router.GetRouter()
	r.Run(":8080")
}
