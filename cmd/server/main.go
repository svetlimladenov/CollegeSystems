package main

import (
	"github.com/svetlimladenov/collegesystems/pkg/server"
)

func main() {
	app := server.NewApp()

	app.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
