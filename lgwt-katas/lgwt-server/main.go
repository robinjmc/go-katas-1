package main

import (
	"log"
	"net/http"
)

// type InMemoryPlayerStore struct{}

// func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
// 	return 123
// }

// func (i *InMemoryPlayerStore) RecordWin(name string) {}

func main() {
	server := &PlayerServer{NewInMemoryPlayerStore()}

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
