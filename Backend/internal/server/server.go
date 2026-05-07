package server

import (
	"fmt"
	"net/http"
)

func Start() error {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "ok")
	})
	return http.ListenAndServe("localhost:8080", nil)
}
