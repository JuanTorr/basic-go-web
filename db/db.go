package db

import (
	"database/sql"
	"io/ioutil"
	"os"
)

var dbpath = "./database.db"

//GetOrCreateDB creates an sqlite database
func GetOrCreateDB() (db *sql.DB, err error) {
	_, err = os.Stat("./database.db")
	dbNotExists := os.IsNotExist(err)
	db, err = sql.Open("sqlite3", "./database.db")
	if err != nil {
		return
	}
	if dbNotExists {
		var dbSchema []byte
		dbSchema, err = ioutil.ReadFile("db/dbschema.sql")
		if err != nil {
			return
		}
		_, err = db.Exec(string(dbSchema))
	}
	return
}
