package staticFileServer

import (
	"log"
	"net/http"
)

func Init(port string) {
	defineRouting(port)
}

func defineRouting(port string) {
	fs := http.FileServer(http.Dir("../static"))
	go http.Handle("/", fs)

	log.Print("Listening on port: ", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
