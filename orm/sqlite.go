package orm

import (
	"database/sql"
	"log"
	"os"

	_ "modernc.org/sqlite"
)

var (
	db     *sql.DB
	dbName *string
	schema string
)

func SQLite_Initialize(dbName_Param *string) {
	dbName = dbName_Param

	configureDatabaseDirectory()
	configureDatabaseFile()

	db_, err := sql.Open("sqlite", "data/"+*dbName+".db")
	if err != nil {
		log.Fatal(err)
		return
	}

	defer db_.Close()
	db = db_

	loadDatabaseSchema()
	executeDatabaseSchema()
}

func configureDatabaseDirectory() {
	path := "data"
	_ = os.Mkdir(path, os.ModePerm)
}

func configureDatabaseFile() {
	file, err := os.Create("data/" + *dbName + ".db")
	if err != nil {
		log.Fatal(err)
		return
	}

	defer file.Close()
}

func loadDatabaseSchema() {
	bytes, err := os.ReadFile("data/schema.sql")
	if err != nil {
		log.Fatal(err)
		return
	}

	schema = string(bytes)
}

func executeDatabaseSchema() {
	_, err := db.Exec(schema)
	if err != nil {
		log.Fatal(err)
		return
	}
}
