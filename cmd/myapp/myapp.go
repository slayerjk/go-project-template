package main

import (
	"flag"
	"log"

	// change this path for your project
	"template/internal/logging"
)

func main() {
	// flags
	logDir := flag.String("log-dir", "logs", "set custom log dir")
	flag.Parse()

	// logging
	appName := "myapp"

	logFile, err := logging.StartLogging(appName, *logDir, 3)
	if err != nil {
		log.Fatalf("failed to start logging:\n\t%s", err)
	}

	defer logFile.Close()

	// code
	log.Print("string 1")
	log.Print("string 2")

}
