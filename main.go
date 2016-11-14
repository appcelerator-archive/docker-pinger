package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const port = ":3000"

func main() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, req *http.Request) {
		hostname, err := os.Hostname()
		if err != nil {
			panic(err)
		}

		response := fmt.Sprintf("[%s] pong", hostname)
		fmt.Fprintln(w, response)
	})

	fmt.Printf("listening on %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
