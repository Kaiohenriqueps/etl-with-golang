package main

import (
	"etl-with-golang/src/csvwrapper"
	"etl-with-golang/src/sqlwrapper"
	"etl-with-golang/src/utils"
	"log"
	"os"
	"sync"
	"time"
)

func process(chunks [][]utils.MyStruct) {
	for i := 0; i < len(chunks); i++ {
		log.Println("Id: ", i)
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			sqlwrapper.InsertStructs(chunks[i])
			wg.Done()
		}()
		wg.Wait()
	}
}

func main() {
	init := time.Now()
	path := os.Args[1]
	objs := csvwrapper.OpenFileAndCreateStruct(path)
	chunks := utils.ChunkSlice(objs, 5000)
	process(chunks)
	log.Println(time.Now().Sub(init))
}
