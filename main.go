package main

import (
	"etl-with-golang/csvwrapper"
	"fmt"
)

func main() {
	fmt.Println("Hello ETL!")
	csvwrapper.FilterFile("data/example.csv")
}
