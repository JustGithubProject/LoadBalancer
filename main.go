package main

import (
	"sync"
)

func main() {
	var wg sync.WaitGroup
	ports := []string{":8001", ":8002", ":8003"}

	for _, port := range ports {
		wg.Add(1)
		go startServer(port, &wg)
	}

	wg.Wait()
}