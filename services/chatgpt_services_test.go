// services/chatgpt_services_test.go
package services

import (
	"os"
	"testing"

	"github.com/baharoam/openaiIntegration/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)


type MockChatGPT struct {
	mock.Mock
}

func (m *MockChatGPT) CallChatGPT(inputs []string) ([]models.LaptopSpec, error) {
	args := m.Called(inputs)
	return args.Get(0).([]models.LaptopSpec), args.Error(1)
}

func TestCallChatGPT(t *testing.T) {
	os.Setenv("OPENAI_API_KEY", "dummy_key") 

	inputs := []string{
		"Laptop: Dell Inspiron; Processor i7-10510U; RAM 16GB; 512GB SSD Missing battery",
	}
	expectedSpecs := []models.LaptopSpec{
		{
			Brand:          "Dell",
			Model:          "Inspiron",
			Processor:      "i7-10510U",
			RamCapacity:    "16GB",
			RamType:       "DDR4", 
			StorageCapacity: "512GB SSD",
			BatteryStatus:  "No",
		},
	}

	mockChatGPT := new(MockChatGPT)
	mockChatGPT.On("CallChatGPT", inputs).Return(expectedSpecs, nil)

	laptopDetails, err := mockChatGPT.CallChatGPT(inputs)

	assert.NoError(t, err)
	assert.Equal(t, expectedSpecs, laptopDetails)

	mockChatGPT.AssertExpectations(t)
}
