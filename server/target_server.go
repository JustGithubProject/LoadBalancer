package main

import (
	"fmt"
	"log"
	"net/http"
)


func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Data from %s", r.Host)
	log.Println(r.Host)
}


func startServer(port string) {
	/*
		Universal function to start server
		port -> Should look like this: ":8001", ":8002"
	*/
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)

	http.ListenAndServe(port, mux)
}	

