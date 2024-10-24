package main

import (
	"fmt"
	"log"

	"github.com/baharoam/openaiIntegration/controllers"
	"github.com/baharoam/openaiIntegration/services"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	// "github.com/baharoam/openaiIntegration/controllers"
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

	r := gin.Default()
	r.POST("/process-laptop", func(c *gin.Context) {
		controllers.ProcessLaptopSpec(c, services.CallChatGPT, "./input/laptops_spec.txt")
	})
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
