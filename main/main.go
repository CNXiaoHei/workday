package main

import (
	"log"
	"net/http"
	"workday"
)

func main() {
	server := workday.NewHTTPServer()
	err := http.ListenAndServe(":8000", server)
	if err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
