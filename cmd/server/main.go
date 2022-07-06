package main

import (
	"net/http"

	"github.com/svetlimladenov/collegesystems/routers"
)

func main() {
	endpoint := "localhost:8080"

	router := routers.InitRouter()

	server := http.Server{
		Addr:    endpoint,
		Handler: router,
	}

	server.ListenAndServe()
}
