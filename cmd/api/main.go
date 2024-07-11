package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/misterpuffin/go-rest-api-boilerplate/internal/config"
	"github.com/misterpuffin/go-rest-api-boilerplate/internal/http"
)

func main() {
	env := flag.String("e", "dev", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()

	config, err := config.LoadConfig(*env)
	if err != nil {
		log.Fatalf("Failed to load config in %s", *env)
	}

	http.RunServer(config)

}
