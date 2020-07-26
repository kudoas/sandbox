package main

import (
	"log"
	"mime"
	"net/http"
	"time"

	"github.com/justinas/alice"
)

func middlewareOne(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Executing middlewareOne")
		next.ServeHTTP(w, r)
		log.Println("Executing middlewareOne again")
	})
}

func middlewareTwo(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Executing middlewareTwo")
		if r.URL.Path == "/foo" {
			return
		}
		next.ServeHTTP(w, r)
		log.Println("Executing middlewareTwo again")
	})
}

func enforceJSONHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		contentType := r.Header.Get("Content-Type")

		if contentType != "" {
			mt, _, err := mime.ParseMediaType(contentType)
			if err != nil {
				http.Error(w, "Malformed Content-Type header", http.StatusBadRequest)
				return
			}
			if mt != "application/json" {
				http.Error(w, "Content-Type header must be application/json", http.StatusUnsupportedMediaType)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}

func final(w http.ResponseWriter, r *http.Request) {
	log.Print("Executing finalHandler")
	w.Write([]byte("OK"))
}

// Smart
func timeHandler(format string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tm := time.Now().Format(format)
		w.Write([]byte("The time is: " + tm))
		log.Print("timeHandler")
	})
}

// Not smart
// type timeHandler struct {
// 	format string
// }

// func (th *timeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	tm := time.Now().Format(th.format)
// 	w.Write([]byte("The time is: " + tm))
// 	log.Println(th)
// }

// th := &timeHandler{time.RFC1123}

func main() {
	// multiplexor
	mux := http.NewServeMux()

	// Chain
	stdChain := alice.New(enforceJSONHandler)
	mdwChain := alice.New(middlewareOne, middlewareTwo)

	// Handler
	rh := http.RedirectHandler("http://example.org", 307)
	th := timeHandler(time.RFC1123)
	finalHandler := http.HandlerFunc(final)

	mux.Handle("/foo", rh)
	mux.Handle("/time", mdwChain.Then(th))
	mux.Handle("/", stdChain.Then(finalHandler))

	log.Println("Listening...")
	http.ListenAndServe(":3000", mux)
}
