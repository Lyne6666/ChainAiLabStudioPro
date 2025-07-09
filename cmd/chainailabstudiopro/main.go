// cmd/chainailabstudiopro/main.go
package main

import (
	"flag"
	"log"
	"os"
	
	"chainailabstudiopro/internal/chainailabstudiopro"
)

func main() {
	verbose := flag.Bool("verbose", false, "Enable verbose logging")
	flag.Parse()
	
	app := chainailabstudiopro.NewApp(*verbose)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
