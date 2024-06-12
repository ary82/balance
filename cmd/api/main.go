package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ary82/balance/internal/infra"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	classifyServerAddr := os.Getenv("CLASSIFY_SERVER_ADDR")
	_ = classifyServerAddr

	port := os.Getenv("PORT")
	dburl := os.Getenv("DB_URL")

	db, err := infra.NewSQLDB(dburl)
	if err != nil {
		log.Fatal(err)
	}

	server := infra.NewFiberServer(db)

	err = server.Listen(fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatal(err)
	}
}
