package main

import (
	"go-app-server/webServer"
	"net/http"
)

var port = "37561"
var staticFileDirectory = http.Dir("./static")

func main() {

	webServer.Initialize(port, staticFileDirectory)
}
