package main

import (
	"os"
	"log"
)

func main() {
	port := string(os.Args[1])
	log.Println(port)
	startServer(":" + port)
}