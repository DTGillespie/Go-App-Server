package dcm

import (
	"database/sql"
	"log"
	"os"

	_ "modernc.org/sqlite"
)

var (
	dbName *string
)

func SQLite_CreateContext(dbName_Param *string) {
	dbName = dbName_Param

	configureDatabaseDirectory()
	configureDatabaseFile()

	db_, err := sql.Open("sqlite", properties.SQLDirectory+*dbName+".db")
	if err != nil {
		log.Fatal(err)
	}

	db = db_
	//defer db_.Close()
}

func configureDatabaseDirectory() {
	path := "data"
	_ = os.Mkdir(path, os.ModePerm)
}

func configureDatabaseFile() {
	file, err := os.Create(properties.SQLDirectory + *dbName + ".db")
	if err != nil {
		log.Fatal(err)
		return
	}

	defer file.Close()
}
