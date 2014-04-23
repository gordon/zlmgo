package main

import (
	"github.com/gordon/zlmgo"
	"log"
	"os"
)

func main() {
	l := zlm.LicenseNew()
	if err := l.Get("My Product", "1.0", os.Args[0], ".", ""); err != nil {
		log.Fatal(err)
	}
}
