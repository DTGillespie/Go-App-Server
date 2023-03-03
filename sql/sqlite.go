package sql

import (
	"database/sql"
	"log"
	"os"

	_ "modernc.org/sqlite"
)

func SQLite_Initialize() {

	configureDirectory()
	configureDatabaseFile()

	db, err := sql.Open("sqlite", "data/test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	schema_Ptr := defineSchemaArr_Ptr()
	schema := *schema_Ptr

	for i := 0; i < len(schema); i++ {
		_, err = db.Exec(schema[i])
		if err != nil {
			log.Fatal(err)
			return
		}
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

func defineSchemaArr_Ptr() *[1]string {

	arr := [1]string{

		` CREATE TABLE inventory_instances(
			ID INTEGER PRIMARY KEY AUTOINCREMENT,
			Name VARCHAR(255) NOT NULL UNIQUE,
			Desc VARCHAR(255)
		);`,
	}

	return &arr
}
