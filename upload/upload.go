package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Import API KEY from .env file
	API_KEY := os.Getenv("API_KEY")

	// Load file
	data, err := ioutil.ReadFile("record1.m4a")

	if err != nil {
		log.Fatal(err)
	}

	// Setup HTTP client and set header
	client := http.Client{}
}