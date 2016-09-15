package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const defPort = "3000"

var port = os.Getenv("PORT")

func main() {
	if port == "" {
		port = defPort
	}

	addr := fmt.Sprintf(":%s", port)

	http.HandleFunc("/ping", pingHandler)

	fmt.Printf("listening on %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func pingHandler(w http.ResponseWriter, req *http.Request) {
	log.Println(req.URL)
	fmt.Fprintln(w, "pong")
}
