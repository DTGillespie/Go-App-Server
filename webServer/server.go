package webServer

import (
	"go-app-server/dcm"
	"log"
	"net/http"
)

type ServerConfigurationSettings struct {
	Port                string
	StaticFileDirectory http.Dir
	DatabaseName        string
	DatabaseDialect     string
}

var (
	settings *ServerConfigurationSettings
	dbConfig dcm.DatabaseProperties
)

func Initialize(settings_Param *ServerConfigurationSettings) {
	settings = settings_Param

	dbConfig = dcm.DatabaseProperties{
		ConnectionHandle: settings.DatabaseName,
		Dialect:          settings_Param.DatabaseDialect,
		SQLDirectory:     "data/",
	}

	dcm.InitializeDatabaseContext(&dbConfig)
	dcm.LoadSchemaFromFile("schema.sql")
	dcm.ExecuteDatabaseSchema()

	defineRouting()

	log.Print("Server Listening on port: ", settings.Port)
	err := http.ListenAndServe(":"+settings.Port, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func defineRouting() {
	route_Root()
}

func route_Root() {
	fs := http.FileServer(settings.StaticFileDirectory)
	http.Handle("/", fs)
}

func route_Request() {

}
