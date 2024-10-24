
package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"

	"github.com/baharoam/openaiIntegration/models"
)

var laptopCache = make(map[string]models.LaptopSpec) 

func CallChatGPT(inputs []string) ([]models.LaptopSpec, error) {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("API Key not found.")
	}

	var laptopDetails []models.LaptopSpec

	for _, input := range inputs {
		modelName := extractModelName(input)

		if cachedSpec, found := laptopCache[modelName]; found {
			laptopDetails = append(laptopDetails, cachedSpec)
			continue
		}

		prompt := createPrompt(input)

		laptopDetail, err := sendChatGPTRequest(apiKey, prompt)
		if err != nil {
			return nil, err
		}

		laptopDetails = append(laptopDetails, laptopDetail)
		laptopCache[modelName] = laptopDetail
	}

	return laptopDetails, nil
}


func createPrompt(input string) string {
	return fmt.Sprintf(`
		Here is some laptop specification data: "%s".
Please convert this data into a structured JSON object with the following format:
{
	"Brand": "",
	"Model": "",
	"Processor": "",
	"RamCapacity": "",
	"RamType": "",
	"StorageCapacity": "",
	"BatteryStatus": ""
}
- Brand refers to the laptop's brand.
- Model refers to the specific model name or number.
- Processor should include the full name of the processor.
- RamCapacity is the size of RAM in GB.
- RamType should be inferred if not explicitly mentioned (e.g., DDR4).
- StorageCapacity is the amount of storage in GB or TB.
- BatteryStatus should be 'Yes' if the battery is included, and 'No' if the battery is missing or damaged.
Please return the result in JSON format.
If any information of laptops is missing or unclear from the provided text, look for this information online to complete the specification
	`, input)
}

func sendChatGPTRequest(apiKey, prompt string) (models.LaptopSpec, error) {
	payload := map[string]interface{}{
		"model":      "gpt-4o-mini",
		"messages":   []map[string]string{{"role": "user", "content": prompt}},
		"max_tokens": 200,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return models.LaptopSpec{}, fmt.Errorf("failed to marshal payload: %v", err)
	}

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(payloadBytes))
	if err != nil {
		return models.LaptopSpec{}, fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return models.LaptopSpec{}, fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return models.LaptopSpec{}, fmt.Errorf("received error %d %s", resp.StatusCode, body)
	}

	var response map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return models.LaptopSpec{}, fmt.Errorf("failed to decode response: %v", err)
	}

	return parseResponse(response)
}


func parseResponse(response map[string]interface{}) (models.LaptopSpec, error) {
	if choices, ok := response["choices"].([]interface{}); ok && len(choices) > 0 {
		if message, ok := choices[0].(map[string]interface{})["message"].(map[string]interface{}); ok {
			if content, ok := message["content"].(string); ok {
				processedContent := cleanResponse(content)
				if isValidJSON(processedContent) {
					var detail models.LaptopSpec
					if err := json.Unmarshal([]byte(processedContent), &detail); err != nil {
						return models.LaptopSpec{}, fmt.Errorf("failed to unmarshal laptop detail: %v", err)
					}
					return detail, nil
				}
			}
		}
	}
	return models.LaptopSpec{}, fmt.Errorf("invalid response format")
}


func extractModelName(input string) string {
	re := regexp.MustCompile(`(?i)(\b[A-Za-z0-9]+(?:\s+[A-Za-z0-9]+)*\b)`)
	matches := re.FindStringSubmatch(input)
	if len(matches) > 0 {
		return matches[0]
	}
	return "Unknown Model"
}


func cleanResponse(content string) string {
	re := regexp.MustCompile(`\{[^}]*\}`)
	return re.FindString(content)
}


func isValidJSON(data string) bool {
	var js map[string]interface{}
	return json.Unmarshal([]byte(data), &js) == nil
}


