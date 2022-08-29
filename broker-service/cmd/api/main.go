package main

import (
	"fmt"
	"log"
	"net/http"
)

const webPost = "80"

type Config struct{}

func main() {
	app := Config{}

	log.Printf("Starting broker service on port %s\n", webPost)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPost),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
