// Database Context Manager
package dcm

import (
	"database/sql"
	"log"
	"os"
)

type DatabaseProperties struct {
	ConnectionHandle string
	Dialect          string
	SQLDirectory     string
}

var (
	properties DatabaseProperties
	db         *sql.DB
	schema     string
)

func InitializeDatabaseContext(properties_Param *DatabaseProperties) {
	properties = *properties_Param

	configureDialect()
}

func LoadSchemaFromFile(path string) {
	bytes, err := os.ReadFile(properties.SQLDirectory + path)
	if err != nil {
		log.Fatal(err)
		return
	}

	schema = string(bytes)
}

func CloseDatabase() {
	defer db.Close()
}

func ExecuteDatabaseSchema() {
	log.Println("DCM Executing Database Schema: \n\n", schema)

	_, err := db.Exec(schema)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func configureDialect() {

	switch properties.Dialect {

	case "SQLite":
		SQLite_CreateContext(&properties.ConnectionHandle)
		break

	}
}
