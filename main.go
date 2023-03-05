package main

import (
	"go-app-server/webServer"
	"log"
	"net/http"

	"gopkg.in/ini.v1"
)

var (
	settings webServer.ServerConfigurationSettings
)

func main() {

	loadConfigurationSettings()

	webServer.Initialize(&settings)
}

func loadConfigurationSettings() {

	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatal(err)
		return
	}

	settings = webServer.ServerConfigurationSettings{
		Port:                cfg.Section("Server").Key("Port").String(),
		StaticFileDirectory: http.Dir(cfg.Section("Server").Key("StaticFileDirectory").String()),
		DatabaseName:        cfg.Section("Database").Key("DatabaseFileName").String(),
		DatabaseDialect:     cfg.Section("Database").Key("Dialect").String(),
	}
}
