package webServer

import (
	"go-app-server/sql"
	"log"
	"net/http"
)

func Initialize(port string, dir http.Dir) {

	initializeDatabase()
	defineRouting(dir)

	log.Print("Server Listening on port: ", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func initializeDatabase() {
	sql.SQLite_Initialize()
}

func defineRouting(dir http.Dir) {
	route_Root(dir)
}

func route_Root(dir http.Dir) {
	fs := http.FileServer(dir)
	http.Handle("/", fs)
}

func route_Request() {

}
