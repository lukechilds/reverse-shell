package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	handler "github.com/lukechilds/reverse-shell/api"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	http.HandleFunc("/", handler.Handler)
	fmt.Printf("Listening on :%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
