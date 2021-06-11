package main

import (
	"errors"
	"fmt"
	"log"
)

var errItFailed = errors.New("it failed")

func run() (string, error) {
	// return "", errItFailed
	return "", fmt.Errorf("failure x")
}

func main() {
	_, err := run()
	if err != nil {
		log.Fatalf("command failed with: %v", err)
	}
	log.Println("all fine")
}
