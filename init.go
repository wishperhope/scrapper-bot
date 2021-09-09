package main

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// load ENV
func init() {
	err := godotenv.Load()
	if err != nil {
		log.Print("Error Loading .env File will use on server instead")
	}

	botID = os.Getenv("BOT_ID")
	if botID == "" {
		log.Fatal("Cannot Load Bot ID")
	}

	minimumWordEnv := os.Getenv("MINIMUM_CHAR")
	result, err := strconv.Atoi(minimumWordEnv)
	if err == nil {
		minimumWord = result
	}
	paginationEnv := os.Getenv("PAGINATION_NUMBER")
	result, err = strconv.Atoi(paginationEnv)
	if err == nil {
		pagination = result
	}
}
