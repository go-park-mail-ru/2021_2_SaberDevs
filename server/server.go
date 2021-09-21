package server

import (
	"fmt"
	"net/http"
)

func Run() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested kek: %s\n", r.URL.Path)
	})

	http.ListenAndServe(":8080", nil)
}