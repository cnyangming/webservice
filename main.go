package main

import (
	"fmt"
	"github.com/cnyangming/webservice/router"
	"github.com/urfave/negroni"
	"net/http"
	"time"
)

func main() {
	mux := router.InitRouter()

	n := negroni.Classic()
	n.UseHandler(mux)

	server := &http.Server{
		Addr:           ":8080",
		Handler:        n,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Println("Start service: http://127.0.0.1:8080")
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
