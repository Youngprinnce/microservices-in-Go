package main

import (
	"fmt"
	"log"
	"net/http"
)

type Config struct {

}

const webPort  = "80"

func main( ) {
	app := Config{}

	log.Println("starting mail service on port", webPort)

	// Define http server
	srv := &http.Server{
		Addr: fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	// Start the server
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
