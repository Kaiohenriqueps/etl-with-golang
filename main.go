package main

import (
	"etl-with-golang/csvwrapper"
	"fmt"
)

func main() {
	fmt.Println("Hello ETL!")
	csvwrapper.OpenFile("data/example.csv")
}
