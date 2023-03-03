package main

import (
	"go-app-server/webServer"
	"net/http"
)

func main() {

	port := "37561"
	staticFileDirectory := http.Dir("./static")
	webServer.Initialize(port, staticFileDirectory)
}
