package main

import (
	"flag"
	"log"
	"time"

	// change this path for your project
	"template/internal/logging"
	"template/internal/rotatefiles"
	// "template/internal/mailing"
)

// log default path & logs to keep after rotation
const (
	appName           = "MYAPP"
	defaultLogPath    = "logs"
	defaultLogsToKeep = 3
	// mailingFile       = "data/mailing.json"
)

var startTime = time.Now()

func main() {
	// flags
	logDir := flag.String("log-dir", defaultLogPath, "set custom log dir")
	logsToKeep := flag.Int("keep-logs", defaultLogsToKeep, "set number of logs to keep after rotation")
	flag.Parse()

	// logging
	logFile, err := logging.StartLogging(appName, *logDir, 3)
	if err != nil {
		log.Fatalf("failed to start logging:\n\t%s", err)
	}
	defer logFile.Close()

	// starting programm notification
	log.Println("Program Started")

	// main code here

	// count & print estimated time
	endTime := time.Now()
	log.Printf("Program Done\n\tEstimated time is %f seconds", endTime.Sub(startTime).Seconds())

	// close logfile and rotate logs
	logFile.Close()

	if err := rotatefiles.RotateFilesByMtime(*logDir, *logsToKeep); err != nil {
		log.Fatalf("failed to rotate logs:\n\t%s", err)
	}
}
