package drivers

import (
	"database/sql"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type DB struct {
	SQL *sql.DB
}

var con = &DB{}

const maxOpenDBConnection = 10
const maxIdleDBConnection = 5
const maxDbLifeTime = 5 * time.Minute

func ConnectSQL(conString string) (*DB, error) {

	db, err := sql.Open("pgx", conString)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(maxOpenDBConnection)
	db.SetConnMaxIdleTime(maxIdleDBConnection)
	db.SetConnMaxLifetime(maxDbLifeTime)

	con.SQL = db

	return con, nil
}
