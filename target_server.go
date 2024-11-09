package main

import (
	"fmt"
	"log"
	"sync"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Data from %s", r.Host)
	log.Println(r.Host)
}

func startServer(port string, wg *sync.WaitGroup) {
	defer wg.Done()

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)

	http.ListenAndServe(port, mux)
}	

