package sql

import (
	"database/sql"
	"log"
	"os"

	_ "modernc.org/sqlite"
)

func SQLite_Initialize() {

	configureDatabaseDirectory()
	configureDatabaseFile()

	db, err := sql.Open("sqlite", "data/test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	schema := loadDatabaseSchema()
	executeDatabaseSchema(db, schema)
}

func configureDatabaseDirectory() {
	path := "data"
	_ = os.Mkdir(path, os.ModePerm)
}

func configureDatabaseFile() {
	file, err := os.Create("data/test.db")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
}

func loadDatabaseSchema() *string {
	bytes, err := os.ReadFile("data/schema.sql")
	if err != nil {
		log.Fatal(err)
	}

	s := string(bytes)
	return &s
}

func executeDatabaseSchema(db *sql.DB, schema *string) {
	_, err := db.Exec(*schema)
	if err != nil {
		log.Fatal(err)
		return
	}
}
