package main

import (
	"log"
	"net/http"
)

// Hard code in memory store temporarily as we implement
// main.go
type InMemoryPlayerStore struct{}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return 123
}

func main() {
	server := &PlayerServer{&InMemoryPlayerStore{}}
	log.Fatal(http.ListenAndServe(":3333", server))
}
