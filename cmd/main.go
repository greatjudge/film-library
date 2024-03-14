package main

import "net/http"

func DummyHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/actors", DummyHandler)     // Get, Post
	mux.HandleFunc("/actor/{id}", DummyHandler) // Get, Put, Patch, Delete

	mux.HandleFunc("/films", DummyHandler)                // Get, Post
	mux.HandleFunc("/film/{id}", DummyHandler)            // Get, Patch, Put, Delete
	mux.HandleFunc("/film/{id}/{actor_id}", DummyHandler) // Post, Delete
}
