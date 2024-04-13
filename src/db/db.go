package db

import (
	"os"

	database "github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/postgresql"
	"github.com/upper/db/v4/adapter/sqlite"
)

func New() (database.Session, error) {
	setting, err := postgresql.ParseURL(os.Getenv("POSTGRESQL_STRING"))
	if err != nil {
		println(err)
		return nil, err
	}

	sess, err2 := postgresql.Open(setting)
	if err2 != nil {
		println(err2)
		println("Session się psuje")

		return nil, err2
	}

	return sess, nil
}

func TestDB() (database.Session, error) {
	setting := sqlite.ConnectionURL{
		Database: "--memory--",
	}

	return sqlite.Open(setting)
}
