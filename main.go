package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Accept request from ", r.RemoteAddr)
		fmt.Fprint(w, "Hello from go.\n")
		return
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
