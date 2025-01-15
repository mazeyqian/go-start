package main

import (
	"flag"
	"log"
	"net/url"

	"github.com/szyhf/go-excel"
)

func defaultUsage(filePath string, sheetNamer interface{}) {
	conn := excel.NewConnecter()
	err := conn.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// Generate an new reader of a sheet
	rd, err := conn.NewReader(sheetNamer)
	if err != nil {
		panic(err)
	}
	defer rd.Close()

	exampleNameMap := make(map[string]int)
	exampleIdMap := map[string]int{}

	for rd.Next() {
		var s struct {
			URL string `xlsx:"url"`
		}
		// Read a row into a struct.
		err := rd.Read(&s)
		if err != nil {
			panic(err)
		}
		urlStr := s.URL
		uu, err := url.Parse(urlStr)
		if err != nil {
			panic(err)
		}
		exampleName := uu.Query().Get("exampleName")
		if exampleName != "" {
			exampleNameMap[exampleName]++
			log.Printf("exampleName: %s", exampleName)
		}
		exampleId := uu.Query().Get("exampleId")
		if exampleId != "" {
			exampleIdMap[exampleId]++
			log.Printf("exampleId: %s", exampleId)
		}
	}
	log.Printf("exampleNameMap: %+v", exampleNameMap)
	for k, v := range exampleNameMap {
		log.Printf("name: %s, count: %d", k, v)
	}
	log.Printf("exampleIdMap: %+v", exampleIdMap)
	for k, v := range exampleIdMap {
		log.Printf("id: %s, count: %d", k, v)
	}
}

// Example: go run scripts/parse-excel-xlsx/main.go -file="data/example_name.xlsx" -sheet="example_name"
// https://github.com/szyhf/go-excel
func main() {
	// Define command-line flags
	filePath := flag.String("file", "", "xlsx file path")
	sheetNamer := flag.String("sheet", "", "sheet name")
	flag.Parse()
	defaultUsage(*filePath, *sheetNamer)
}
