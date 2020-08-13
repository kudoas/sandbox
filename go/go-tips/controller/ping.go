package controller

import (
	"net/http"

	"github.com/jmoiron/sqlx"
)

type Ping struct {
	db *sqlx.DB
}

func NewPing(db *sqlx.DB) *Ping {
	return &Ping{db: db}
}

func (p *Ping) Index(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	if err := p.db.Ping(); err != nil {
		return http.StatusInternalServerError, nil, err
	}
	return http.StatusOK, "success to connect db server", nil
}
