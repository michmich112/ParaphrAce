package main

import (
	"fmt"
	"log"
	"net/http"
	"server/router"
)

func main() {
	port := 8080
	r := router.Router()

	fmt.Printf("Starting server on port %d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}
