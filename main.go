package main

import (
	"go-app-server/webServer"
	"net/http"
)

var port = "37561"
var staticFileDirectory = http.Dir("./static")

func main() {

	//config := loadConfigurationSettings()

	webServer.Initialize(port, staticFileDirectory)
}

func loadConfigurationSettings() {
	//https:\\ini.unknwon.io/docs/intro/getting_started
}
