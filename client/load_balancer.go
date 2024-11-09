package main

type LoadBalancer struct {
	servers []string
	index int 
}


func (lb *LoadBalancer) getNextServer() string {
	if lb.index == len(lb.servers) {
		lb.index = 0
	}
	currentIndex := lb.index
	lb.index++
	return lb.servers[currentIndex]
}
