package db

import (
	"database/sql"
	"log"
	"os"

	_ "modernc.org/sqlite"

)

const schema = `
CREATE TABLE IF NOT EXISTS scheduler (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	date CHAR(8) NOT NULL DEFAULT "",
	title VARCHAR NOT NULL DEFAULT "",
	comment TEXT DEFAULT "",
	repeat VARCHAR(128) DEFAULT ""
);
CREATE INDEX IF NOT EXISTS idx_date ON scheduler (date);
`

var db *sql.DB

func Init(dbFile string) error {
	var install bool

	_, err := os.Stat(dbFile)
	if err != nil {
		install = true
	}

	db, err = sql.Open("sqlite", dbFile)
	if err != nil {
		log.Println("error in sql.Open function", err)
		return err
	}


	if install == true {
		_, err = db.Exec(schema)
		if err != nil {
			log.Println("error in db.Exec function", err)
			return err
		}
	}
	return nil
}

func Close()  {
	if db != nil {
		db.Close()
	}
	
}