package poll

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

const TRANSCRIPT_URL = "https://api.assemblyai.com/v2/transcript"

func GetTextTranscripted(idTranscript string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Import API KEY from .env file
	API_KEY := os.Getenv("API_KEY")

	// Build POLLING_URL
	pollingUrl := TRANSCRIPT_URL + "/" + idTranscript

	client := &http.Client{}

	req, _ := http.NewRequest("GET", pollingUrl, nil)
	req.Header.Set("content-type", "application/json")
	req.Header.Set("authorization", API_KEY)

	res, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	// Decode
	var result map[string]interface{}
	json.NewDecoder(res.Body).Decode(&result)

	if result["status"] == "completed" {
		return result["text"].(string)
	}

	return ""
}
