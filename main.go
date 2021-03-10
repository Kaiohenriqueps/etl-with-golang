package main

import (
	"etl-with-golang/src/csvwrapper"
	"log"
	"time"
)

func main() {
	init := time.Now()
	csvwrapper.InsertStructs("data/base_teste.txt")
	log.Println(time.Now().Sub(init))
}
