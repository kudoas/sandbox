package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/kudoas/enjoy-middleware/middleware"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func CurrentTimeHandler(w http.ResponseWriter, r *http.Request) {
	curTime := time.Now().Format(time.Kitchen)
	w.Write([]byte(fmt.Sprintf("the current time is %v", curTime)))
}

func middlewareWrapper(f http.HandlerFunc) http.HandlerFunc {
	return middleware.Logger(middleware.Header(f))
}

func main() {
	http.HandleFunc("/v1/hello", middlewareWrapper(HelloHandler))
	http.HandleFunc("/v1/time", middlewareWrapper(CurrentTimeHandler))

	log.Printf("server is listening at %s", ":8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
