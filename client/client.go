package main

import (
	"io"
	"log"
	"net/http"
)

var servers = []string{
	"http://localhost:8001",
	"http://localhost:8002",
	"http://localhost:8003",
}

// Creating object of LoadBalancer
var balancer = LoadBalancer{
	servers: servers,
	index: 0,
}

func sendRequest(path string) {
	client := &http.Client{}
	requestTo := balancer.getNextServer() + path
	log.Println(requestTo)

	response, err := client.Get(requestTo)
	if err != nil {
		log.Fatal(err)
		return
	}

	defer response.Body.Close()

	// Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Println("Response status: ", response.Status)
	log.Println("Response body: ", string(body))
}


func main() {
	for i := 0; i < 6; i++ {
		sendRequest("/")
	}
}