// cmd/secureshield/main.go
package main

import (
	"flag"
	"log"
	"os"

	"secureshield/internal/secureshield"
)

func main() {
	verbose := flag.Bool("verbose", false, "Enable verbose logging")
	flag.Parse()

	app := secureshield.NewApp(*verbose)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
