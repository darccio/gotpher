package main

import (
	"log"
	"os"
)

func main() {
	gotpher := newCliApp()
	err := gotpher.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
