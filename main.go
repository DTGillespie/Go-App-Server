package main

import (
	staticFileServer "go-app-server/modules"
)

func main() {

	moduleInitialization()
}

func moduleInitialization() {
	staticFileServer.Init("37561")
}
