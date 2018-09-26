package db

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

func PsqlConn() *sql.DB {
	var db *sql.DB
	var err error

	for db, err = sql.Open("postgres", psqlConnStr); err != nil; {
		log.Printf("[Postgres Connect] Err: %+v", errors.Wrap(err, ""))
		time.Sleep(time.Duration(psql.errSleepTime) * time.Millisecond)

		db, err = sql.Open("postgres", psqlConnStr)
	}
	log.Println("[Postgres Connect] Port:", psql.port)

	return db
}
