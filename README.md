
## Features

- Reads raw laptop specification data from a text file.
- Connects to the OpenAI GPT API to structure the data into a standardized format.
- Caches processed laptop specs to improve performance. (with in-memory map)
- Exposes an API endpoint (`/process-laptop`) to process the laptop specification data.

## Project Structure

- **`main.go`**: The entry point for the application.
- **`controllers/`**: Contains the `ProcessLaptopSpec` controller that handles reading and processing laptop specifications.
- **`services/`**: Contains the logic for connecting to the OpenAI GPT API.
- **`models/`**: Contains the `LaptopSpec` struct.
- **`input/`**: Contains the text files (`laptops_spec.txt`) with raw laptop specifications.

## Requirements

- Go 1.18+
- OpenAI API Key
- Gin (for the web framework)
- Testify (for testing)
- Git

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/baharoam/openaiIntegration.git
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Set up  OpenAI API key:
   Create a `.env` file in the project root and add OpenAI API key:
   ```bash
   OPENAI_API_KEY=your_openai_api_key
   ```

4. (Optional) Install `godotenv` to load environment variables from `.env` files:
   ```bash
   go get github.com/joho/godotenv
   ```

## How to Run the Project

1. **Ensure your `.env` file is properly set up** with your OpenAI API key.
   
2. **Run the API server**:
   ```bash
   go run main.go
   ```

3. The API server will run on `http://localhost:8080`.

## API Endpoints

### POST `/process-laptop`

This endpoint reads the raw laptop specification data from a file (`./input/laptops_spec.txt`), processes it using the OpenAI API, and returns a structured JSON response.

#### Request:

- Method: `POST`
- URL: `http://localhost:8080/process-laptop`

#### Example Request Body:

None. This API reads the file directly from the `input` folder.

#### Response:

```json
[
  {
    "Brand": "Dell",
    "Model": "Inspiron",
    "Processor": "i7-10510U",
    "RamCapacity": "16GB",
    "RamType": "DDR4",
    "StorageCapacity": "512GB SSD",
    "BatteryStatus": "No"
  },
  {
    "Brand": "Apple",
    "Model": "MacBook Pro",
    "Processor": "M1",
    "RamCapacity": "8GB",
    "RamType": "LPDDR4X",
    "StorageCapacity": "256GB SSD",
    "BatteryStatus": "No"
  }
  ...
]
```

### Input File Example (`./input/laptops_spec.txt`)

The input file should contain dirty, unstructured laptop data:

```
Laptop: Dell Inspiron; Processor i7-10510U; RAM 16GB; 512GB SSD Missing battery
MacBook Pro with M1 chip, 8GB RAM, 256 GB SSD storage Battery removed
ThinkPad, i5 CPU, 8GB memory, storage: 1TB HDD
Asus ROG, Processor: AMD Ryzen 7; RAM 16 GB; 1TB SSD; Damaged battery
```

### Customizing Input File

To customize the input, modify the `laptops_spec.txt` file in the `./input` folder.

## Running Tests

Unit tests are provided for the core functionality (e.g., reading files, processing laptop data, interacting with OpenAI). The tests use `testify` for mocking and assertions.

### Running Tests

1. Run cd services and then Run the tests using the `go test` command:
   ```bash
   go test ./...
   ```

2. Ensure that all tests pass:
   - Mocking the `CallChatGPT` function ensures that we do not make actual API calls during testing.

### Example Test Output

```bash
ok  	github.com/yourusername/openai-laptop-specs/services	0.031s
```


## Caching Behavior

This project uses an in-memory map (`laptopCache`) to cache processed laptop specifications. If the same laptop model is encountered again, it returns the cached response without calling the OpenAI API.
