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
