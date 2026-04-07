package main

import "net/http"

func handler(w http.ResponseWriter, r *http.Request) {
	s := []byte("haLLO")

	if _, err := w.Write(s); err != nil {
		return
	}

}
