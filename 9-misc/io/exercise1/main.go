package main

// Exercise: Write a program, that takes one argument - a filename - an prints
// its content to stdout.

import (
	"flag"
	"io"
	"log"
	"os"
)

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		log.Fatal("file required")
	}
	f, err := os.Open(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	if _, err := io.Copy(os.Stdout, f); err != nil {
		log.Fatal(err)
	}
}
