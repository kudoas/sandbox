package main

import (
	"log"
	"net/http"
	"time"
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

// func final(w http.ResponseWriter, r *http.Request) {
// 	log.Print("Executing finalHandler")
// 	w.Write([]byte("OK"))
// }

func timeHandler(format string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tm := time.Now().Format(format)
		w.Write([]byte("The time is: " + tm))
		log.Print("timeHandler")
	})
}

// 上の方がスマート
// type timeHandler struct {
// 	format string
// }

// func (th *timeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	tm := time.Now().Format(th.format)
// 	w.Write([]byte("The time is: " + tm))
// 	log.Println(th)
// }

func main() {
	// multiplexor
	mux := http.NewServeMux()

	rh := http.RedirectHandler("http://example.org", 307)
	mux.Handle("/foo", rh)

	// th := &timeHandler{time.RFC1123}
	th := timeHandler(time.RFC1123)
	mux.Handle("/time", middlewareOne(middlewareTwo(th)))

	log.Println("Listening...")
	http.ListenAndServe(":3000", mux)
}
