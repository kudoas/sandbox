package dbutil

import (
	"github.com/jmoiron/sqlx"
)

func Transact(db *sqlx.DB, txFunc func(*sqlx.Tx) error) error {
	tx, err := db.Beginx()
	if err != nil {
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()
	if err = txFunc(tx); err != nil {
		return err
	}
	return nil
}
