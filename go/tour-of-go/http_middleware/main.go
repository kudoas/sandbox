package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/justinas/alice"
	"github.com/kudoas/enjoy-middleware/middleware"
	"github.com/kudoas/enjoy-middleware/model"
	"github.com/kudoas/enjoy-middleware/response"
	_ "github.com/lib/pq"
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

func GetUserHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := model.FetchUser(db)
		if err != nil {
			log.Fatalf("failed to fetch user: %v", err)
			response.InternalServerError(w)
			return
		}

		b, err := json.Marshal(users)
		if err != nil {
			log.Fatalf("failed to marshal user: %v", err)
			response.InternalServerError(w)
			return
		}
		_, err = w.Write(b)
		if err != nil {
			log.Fatalf("failed to write response: %v", err)
			response.InternalServerError(w)
			return
		}

		return
	}
}

func main() {
	db, err := sql.Open("postgres", "host=localhost port=5432 dbname=default_db user=postgres password=default_pass sslmode=disable")
	if err != nil {
		log.Fatalf("failed to open db: %v", err)
	}
	defer db.Close()

	logger, err := middleware.NewLogger()
	if err != nil {
		log.Fatalln(err)
	}
	chain := alice.New(middleware.Authentication, middleware.Logger(logger), middleware.Header)

	http.Handle("/hello", chain.Then(http.HandlerFunc(HelloHandler)))
	http.Handle("/user", chain.Then(http.HandlerFunc(GetUserHandler(db))))

	log.Printf("server is listening at %s", ":8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalln(err)
	}
}
