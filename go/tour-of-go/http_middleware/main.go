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

type Response struct {
	Message string `json:"message"`
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	res := Response{
		Message: "hello world",
	}
	b, _ := json.Marshal(res)
	_, _ = w.Write(b)
}

func CurrentTimeHandler(w http.ResponseWriter, r *http.Request) {
	curTime := time.Now().Format(time.Kitchen)
	res := Response{
		Message: fmt.Sprintf("the current time is %v", curTime),
	}
	b, _ := json.Marshal(res)
	_, _ = w.Write(b)
}

func main() {
	logger, err := middleware.NewLogger()
	if err != nil {
		log.Fatalln(err)
	}
	chain := alice.New(middleware.Authentication, middleware.Logger(logger), middleware.Header)

	http.Handle("/v1/hello", chain.Then(http.HandlerFunc(HelloHandler)))
	http.Handle("/v1/time", chain.Then(http.HandlerFunc(CurrentTimeHandler)))

	log.Printf("server is listening at %s", ":8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalln(err)
	}
}
