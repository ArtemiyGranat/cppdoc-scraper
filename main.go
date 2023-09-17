package main

import (
	"cppreference-scraper/scraper"
	"cppreference-scraper/config"

	"encoding/json"
	"fmt"
	"os"
)

func writeJSONToFile(fileName string, data []byte) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	domains := []string{"ru.cppreference.com"}
	url := "https://ru.cppreference.com/w/cpp"
	cacheDir := "./cpp-reference-cache"
	config := config.New(domains, url, cacheDir)

	scraper := scraper.New(config)
	
	pages := scraper.Scrape()
	jsonData, err := json.MarshalIndent(pages, "", "    ")
	if err != nil {
		panic(err)
	}

	fileName := "pages.json"
	err = writeJSONToFile(fileName, jsonData)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Data has been written to %s\n", fileName)
}