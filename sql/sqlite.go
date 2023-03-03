package sql

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func SQLite_Initialize() {

	configureDirectory()
	configureDatabaseFile()

	db, err := sql.Open("sqlite3", "data/test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	testStatement := `
	create table foo (id integer not null primary key, name text);
	delete from foo;
	`

	_, err = db.Exec(testStatement)
	if err != nil {
		log.Fatal(err)
		return
	}

}

func configureDirectory() {
	path := "data"
	_ = os.Mkdir(path, os.ModePerm)
}

func configureDatabaseFile() {
	f, err := os.Create("data/test.db")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
}
