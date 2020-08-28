package main

import (
	"log"
	"net/http"

	"github.com/srehouni/api-go/pkg/server"
)

func main() {

	s := server.New()
	log.Fatal(http.ListenAndServe(":8082", s.Router()))
}
