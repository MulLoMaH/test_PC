package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hallo")

	http.HandleFunc("/res", handler)

	fmt.Println("Server start")
	http.ListenAndServe(":8080", nil)

}

func handler(w http.ResponseWriter, r *http.Request) {
	s := []byte("haLLO")

	if _, err := w.Write(s); err != nil {
		return
	}

}
