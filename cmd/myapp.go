package main

import (
	"flag"
	"log"
	"time"

	// change this path for your project
	"github.com/slayerjk/go-project-template/internal/logging"
	"github.com/slayerjk/go-project-template/internal/vafswork"
	// "github.com/slayerjk/go-project-template/internal/mailing"
)

func main() {
	// defining default values
	var (
		appNameDefault string    = "MY-APP"
		startTime      time.Time = time.Now()
		// mailingFile       string = getExePath() + "/data/mailing.json"
	)

	// flags
	appName := flag.String("app-name", appNameDefault, "set application name(used for logs name, mailing subject, etc)")

	var LogsPath string = vafswork.GetExePath() + "/logs" + "_" + *appName
	logsDir := flag.String("log-dir", LogsPath, "set custom log dir")

	logsToKeep := flag.Int("keep-logs", 7, "set number of logs to keep after rotation")
	flag.Parse()

	// logging
	logFile, err := logging.StartLogging(*appName, *logsDir, *logsToKeep)
	if err != nil {
		log.Fatalf("failed to start logging:\n\t%s", err)
	}
	defer logFile.Close()

	// starting programm notification
	log.Printf("%s Started", *appName)

	// main code here

	// count & print estimated time
	endTime := time.Now()
	log.Printf("Program Done\n\tEstimated time is %f seconds", endTime.Sub(startTime).Seconds())

	// close logfile and rotate logs
	logFile.Close()

	if err := vafswork.RotateFilesByMtime(*logsDir, *logsToKeep); err != nil {
		log.Fatalf("failed to rotate logs:\n\t%s", err)
	}
}
