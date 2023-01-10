package main

import (
	"github.com/anthonysyk/fizzbuzz/api"
	"github.com/anthonysyk/fizzbuzz/store"
	"log"
	"net/http"
)

func main() {
	st := store.NewMemoryStore()
	router := api.GetRouter(st)
	log.Fatal(http.ListenAndServe(":8080", router))
}
