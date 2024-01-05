package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/justinas/alice"
	"github.com/kudoas/enjoy-middleware/middleware"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	b, _ := json.Marshal(`{"message": "Hello, World!"}`)
	w.Write(b)
}

func CurrentTimeHandler(w http.ResponseWriter, r *http.Request) {
	curTime := time.Now().Format(time.Kitchen)
	msg := fmt.Sprintf(`{"message": "the current time is %v"}`, curTime)
	b, _ := json.Marshal(msg)
	w.Write(b)
}

func main() {
	chain := alice.New(middleware.Authentication, middleware.Logger, middleware.Header)

	http.Handle("/v1/hello", chain.Then(http.HandlerFunc(HelloHandler)))
	http.Handle("/v1/time", chain.Then(http.HandlerFunc(CurrentTimeHandler)))

	log.Printf("server is listening at %s", ":8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalln(err)
	}
}
