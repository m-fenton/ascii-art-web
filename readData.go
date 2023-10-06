package main

import (
	"log"
	"os"
)

func readData(banner string) string {
	data, err := os.ReadFile(banner)
	if err != nil {
		log.Fatal(err)
	}

	return string(data)
}
