package controllers

import (
	"bufio"
	"log"
	"net/http"
	"os"
	"github.com/gin-gonic/gin"
)

var inputs []string
func ReadLaptopSpecFromFile(c *gin.Context){
	filePath := "input/laptops_spec.txt"
	file, err := os.Open(filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not open input file"})
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputs = append(inputs, scanner.Text())
	}
	log.Println(inputs[0])

	if err := scanner.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading input file"})
		return
	}
}


func ProcessLaptopSpec(c *gin.Context){
	ReadLaptopSpecFromFile(c)
//	CallOpenaiService(c)
}