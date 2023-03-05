package webServer

import (
	"go-app-server/orm"
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
)

func Initialize(settings_Param *ServerConfigurationSettings) {
	settings = settings_Param

	initializeDatabaseContext()
	defineRouting()

	log.Print("Server Listening on port: ", settings.Port)
	err := http.ListenAndServe(":"+settings.Port, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func initializeDatabaseContext() {

	switch settings.DatabaseDialect {

	case "SQLite":
		orm.SQLite_Initialize(&settings.DatabaseName)
		break

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
