package main

import (
	"log"
	"os"

	"github.com/kelseyhightower/journal-2-logentries/journal"
	"github.com/kelseyhightower/journal-2-logentries/logentries"
)

func main() {
        log.Print("Logging to the console works")
	socket := os.Getenv("LOGENTRIES_JOURNAL_SOCKET")
	if socket == "" {
		socket = journal.DefaultSocket
	}
        log.Print(socket)
	url := os.Getenv("LOGENTRIES_URL")
	if url == "" {
		url = logentries.DefaultUrl
	}
        log.Print(url)
	token := os.Getenv("LOGENTRIES_TOKEN")
	if token == "" {
		log.Fatal("non-empty input token (LOGENTRIES_TOKEN) is required. See https://logentries.com/doc/input-token")
	}
        log.Print(token)
	logs, err := journal.Follow(socket)
	if err != nil {
		log.Fatal(err.Error())
	}
	le, err := logentries.New(url, token)
	if err != nil {
		log.Fatal(err.Error())
	}

        log.Print("Starting journal-2-logentries...")

	for {
		select {
		case data := <-logs:
			if _, err := le.Write(data); err != nil {
				log.Print(err.Error())
			}
		}
	}
}
