package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	if flag.NArg() == 0 {
		log.Fatal("URL required")
	}
	resp, err := http.Get(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
}
