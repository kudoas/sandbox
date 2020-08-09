package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// handler: ServeHTTPというメソッドを持ったインターフェース
type AppHandler struct {
	h interface{}
}

func (a AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world")
}

// goでのjsonの取り扱い
type Message struct {
	Name string
	Time int
}

// handler function: handlerのように振る舞う関数
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	m := Message{"kudoa", 122}
	// encode json data
	b, _ := json.Marshal(m)
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	// handler: middlewareを噛ませやすい？
	http.Handle("/", AppHandler{})
	http.HandleFunc("/hello", HelloHandler)

	server.ListenAndServe()
}
