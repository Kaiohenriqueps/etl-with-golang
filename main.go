package main

import (
	"etl-with-golang/src/csvwrapper"
	"etl-with-golang/src/sqlwrapper"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	init := time.Now()
	path := os.Args[1]
	objs := csvwrapper.OpenFileAndCreateStruct(path)
	sqlwrapper.InsertStructs(objs)
	log.Println(time.Now().Sub(init))
}
