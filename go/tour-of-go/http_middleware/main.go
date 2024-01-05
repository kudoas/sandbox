package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type Logger struct {
	m *http.ServeMux
}

func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	l.m.ServeHTTP(w, r)
	log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
}

func NewLogger(m *http.ServeMux) *Logger { return &Logger{m} }

type ResponseHeader struct {
	m           *http.ServeMux
	headerName  string
	headerValue string
}

func (rh *ResponseHeader) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add(rh.headerName, rh.headerValue)
	log.Println(rh.headerName, rh.headerValue)
	rh.m.ServeHTTP(w, r)
}

func NewResponseHeader(m *http.ServeMux, name string, value string) *ResponseHeader {
	return &ResponseHeader{m: m, headerName: name, headerValue: value}
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func CurrentTimeHandler(w http.ResponseWriter, r *http.Request) {
	curTime := time.Now().Format(time.Kitchen)
	w.Write([]byte(fmt.Sprintf("the current time is %v", curTime)))
}

func main() {
	http.HandleFunc("/v1/hello", HelloHandler)
	http.HandleFunc("/v1/time", CurrentTimeHandler)

	wrapMux := NewLogger(NewResponseHeader(http.DefaultServeMux, "X-My-Header", "my header value").m)

	log.Printf("server is listening at %s", ":8080")
	log.Fatal(http.ListenAndServe(":8080", wrapMux))
}
