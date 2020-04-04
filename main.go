package main

import (
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
	"net/http"
	"time"
)

var formatter = render.New()

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//fmt.Println("Accept request from ", r.RemoteAddr)
		formatter.Text(w, http.StatusOK, "Hello from go")
		return
	})
	mux.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		//fmt.Println("Accept request from ", r.RemoteAddr)
		formatter.JSON(w, http.StatusOK, struct {
			Author string
		}{"yang ming"})
		return
	})

	n := negroni.Classic()
	n.UseHandler(mux)

	server := &http.Server{
		Addr:           ":8080",
		Handler:        n,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
