package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
//	"github.com/baharoam/openaiIntegration/controllers"
)

func loadEnv() {
    err := godotenv.Load(".env")
    if err != nil {
        log.Fatal("Error loading .env file")
    }
    fmt.Println("Environment variables loaded.")
}


func main() {
	loadEnv()
	//api_key := os.Getenv("OPENAI_API_KEY")

	r := gin.Default()
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
