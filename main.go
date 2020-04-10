package main

import (
	"context"
	"fmt"
	"github.com/cnyangming/webservice/router"
	"github.com/urfave/negroni"
	"log"
	"net/http"
	"os"
	"os/signal"
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

	go func() {
		fmt.Println("Start service: http://127.0.0.1:8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalln("Listen: ", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("ShutDown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Println("Server ShutDown: ", err)
	}
	log.Println("Server Exiting.")
}
