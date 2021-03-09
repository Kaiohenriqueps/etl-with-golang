package main

import (
	"etl-with-golang/csvwrapper"
	"fmt"
)

func main() {
	fmt.Println("Hello ETL!")
	csvwrapper.FilterFile("data/input/example.csv")
	// sqlwrapper.ConnectToPostgres()
}
