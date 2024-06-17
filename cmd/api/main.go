package main

import (
	"log"
	"os"

	"github.com/ary82/balance/internal/infra"

	"github.com/joho/godotenv"
)

func main() {
	// Load envvars based on mode
	mode := os.Getenv("MODE")
	if mode != "prod" {
		err := godotenv.Load("./.env")
		if err != nil {
			log.Fatal(err)
		}
	}

	classifyServerAddr := os.Getenv("CLASSIFY_SERVER_ADDR")
	port := os.Getenv("PORT")
	dburl := os.Getenv("DB_URL")

	if classifyServerAddr == "" {
		log.Fatal("classifyServerAddr envvar empty")
	}
	if port == "" {
		log.Fatal("port envvar empty")
	}
	if dburl == "" {
		log.Fatal("dburl envvar empty")
	}

	err := infra.Run(dburl, classifyServerAddr, port)
	if err != nil {
		log.Fatal(err)
	}
}
