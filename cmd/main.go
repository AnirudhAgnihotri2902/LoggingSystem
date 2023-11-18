package main

import (
	elasticsearch "github.com/elastic/go-elasticsearch/v8"
	"log"
	"os"
)

const (
	// exit is exit code which is returned by realMain function.
	// exit code is passed to os.Exit function.
	exitOK int = iota
	exitError
)

func main() {
	os.Exit(realMain(os.Args))
}

func realMain(args []string) int {
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}
	log.Println(elasticsearch.Version)
	res, err := es.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()
	return exitOK
}
