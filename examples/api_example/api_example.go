package main

import (
	"log"
	"os"

	"github.com/wikena/zlmgo"
)

func main() {
	l := zlmgo.LicenseNew()
	if err := l.Get("My Product", "1.0", os.Args[0], ".", ""); err != nil {
		log.Fatal(err)
	}
}
