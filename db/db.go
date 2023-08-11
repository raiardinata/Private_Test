package db

import (
	"log"

	"private_test/config"

	"github.com/gocraft/dbr"

	_ "github.com/lib/pq"
)

var (
	db    *dbr.Connection
	err   error
	debug string = ""
)

// Init the database
func Init() {
	conf := config.GetConfig()
	connStr := "host=" + conf.DB_HOST + " port=" + conf.DB_PORT + " user=" + conf.DB_USERNAME + " " + "password=" + conf.DB_PASSWORD + " dbname=" + conf.DB_NAME + " sslmode=disable"

	db, err = dbr.Open("postgres", connStr, nil)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err.Error() + " DSN error")
	}
}

func CreatCon() *dbr.Connection {
	return db
}

// PgNewSession Creates a new session
func PgNewSession() *dbr.Session {
	if debug == "true" {
		return db.NewSession(DbrLogger)
	}
	return db.NewSession(nil)
}
