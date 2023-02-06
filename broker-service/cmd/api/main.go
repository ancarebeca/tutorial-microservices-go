package main

import (
	"fmt"
	"log"
	"net/http"
)

const webPort = "80"

func main() {
	app := Config{}

	log.Printf("Starting broken on port %s\n", webPort)

	// define http server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	// start the server
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
