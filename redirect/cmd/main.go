package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		h := r.Host

		if r.Method == "GET" {

			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "Host: "+h)

			return

		}

	})

	fmt.Println("Listening on port 4000...")

	http.ListenAndServe(":4000", nil)
}
