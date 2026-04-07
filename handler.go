package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 3; i++ {
			time.Sleep(1 * time.Second)
			if err := json.NewEncoder(w).Encode(map[string]string{"Hallo": "World"}); err != nil {
				log.Println("error", err)
			}
		}
	}()

	wg.Wait()
}

func handlerINT(w http.ResponseWriter, r *http.Request) {

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 3; i++ {
			time.Sleep(1 * time.Second)
			if err := json.NewEncoder(w).Encode(map[string]int{"Hallo": i}); err != nil {
				log.Println("error", err)
			}
		}
	}()

	wg.Wait()
}
