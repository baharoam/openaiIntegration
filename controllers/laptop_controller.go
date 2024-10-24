package controllers

import (
	"bufio"
	"log"
	"net/http"
	"os"

	"github.com/baharoam/openaiIntegration/models"
	"github.com/baharoam/openaiIntegration/services"
	"github.com/gin-gonic/gin"
)

func ReadLaptopSpecFromFile(c *gin.Context, filePath string) ([]string, error) {
    file, err := os.Open(filePath)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not open input file"})
        return nil, err
    }
    defer file.Close()

    var inputs []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        inputs = append(inputs, scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading input file"})
        return nil, err
    }

    log.Printf("Read lines: %v", inputs)
    return inputs, nil
}


func CallOpenaiService(c *gin.Context){
	filePath := "input/laptops_spec.txt"
	inputs, err := ReadLaptopSpecFromFile(c, filePath)
	if err != nil {
		return
	}

	laptopDetails, err := services.CallChatGPT(inputs)
	if err != nil {
	 c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	 return
	}
	c.JSON(http.StatusOK, laptopDetails)
}

func ProcessLaptopSpec(c *gin.Context, callChatGPT func([]string) ([]models.LaptopSpec, error), filePath string) {
    data, err := ReadLaptopSpecFromFile(c, filePath)
    if err != nil || len(data) == 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
        return
    }

    laptopSpecs, err := callChatGPT(data)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to call ChatGPT"})
        return
    }

    c.JSON(http.StatusOK, laptopSpecs)
}
