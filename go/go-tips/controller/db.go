package controller

import (
	"net/http"

	"github.com/jmoiron/sqlx"
)

func PingDB(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	dbcon, err := sqlx.Open("sqlite3", "blog.sqlite3")
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	defer dbcon.Close()
	if err := dbcon.Ping(); err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, "success to connect db server", nil
}
