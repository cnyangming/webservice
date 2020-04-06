package router

import (
	"github.com/unrolled/render"
	"log"
	"net/http"
	"os"
)

var formatter = render.New()

func InitRouter() *http.ServeMux {
	webRoot := getWebRoot() + "/assets"
	mux := http.NewServeMux()

	mux.Handle("/", http.FileServer(http.Dir(webRoot)))
	mux.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		//fmt.Println("Accept request from ", r.RemoteAddr)
		formatter.JSON(w, http.StatusOK, struct {
			Author string
		}{"yang ming"})
		return
	})

	return mux
}

func getWebRoot() string {
	webRoot := os.Getenv("WEBROOT")
	if len(webRoot) == 0 {
		root, err := os.Getwd()
		if err != nil {
			log.Fatalf("getWebroot fial %v\n", err)
		}
		return root
	}
	return webRoot
}
