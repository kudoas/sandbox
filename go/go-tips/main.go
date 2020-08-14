package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"example.com/user/go-tips/controller"
	"example.com/user/go-tips/db"
	"example.com/user/go-tips/repository"
)

// handler: ServeHTTPというメソッドを持ったインターフェース
type AppHandler struct {
	// コントローラーを定義
	h func(http.ResponseWriter, *http.Request) (int, interface{}, error)
}

func (a AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	code, res, err := a.h(w, r)
	if err != nil {
		w.WriteHeader(code)
		_, writeErr := w.Write([]byte(err.Error()))
		if writeErr != nil {
			log.Print(writeErr)
		}
		return
	}
	w.WriteHeader(code)

	response, err := json.Marshal(res)
	if _, writeErr := w.Write(response); writeErr != nil {
		log.Print(writeErr)
	}
	return
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

// HelloHandler := HandlerFunc(Hello)
// HandleFunc: ハンドル関数をDefaultServerMuxに登録する関数
// HandlerFunc: 型の名前

// handler or handler func
// 設計の問題：既存のインターフェース、もしくはハンドラとして使える型が欲しいなら、
// 単にそのインターフェースにメソッドServeHTTP追加すれば、URLに割り当てられるハンドラを得られる

func main() {
	cs := db.NewDB("blog.sqlite3")
	dbcon, err := cs.Open()
	if err != nil {
		log.Fatal(err)
	}
	if err := repository.CreateTable(dbcon); err != nil {
		log.Fatal(err)
	}

	postController := controller.NewPost(dbcon)
	pingController := controller.NewPing(dbcon)

	r := mux.NewRouter()
	r.Methods(http.MethodGet).Path("/post/{id:[0-9]+}").Handler(AppHandler{postController.Show})
	r.Methods(http.MethodGet).Path("/post").Handler(AppHandler{postController.Index})
	r.Methods(http.MethodPost).Path("/post").Handler(AppHandler{postController.Create})
	r.Methods(http.MethodPut).Path("/post/{id:[0-9]+}").Handler(AppHandler{postController.Update})

	r.Methods(http.MethodGet).Path("/ping").Handler(AppHandler{pingController.Index})

	if err := http.ListenAndServe("127.0.0.1:8080", handlers.CombinedLoggingHandler(os.Stdout, r)); err != nil {
		log.Fatal(err)
	}
}
