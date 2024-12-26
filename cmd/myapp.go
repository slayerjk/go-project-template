package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	// change this path for your project
	logging "github.com/slayerjk/go-logging"
	vafswork "github.com/slayerjk/go-vafswork"
	// mailing "github.com/slayerjk/go-mailing"
	// vawebwork "github.com/slayerjk/go-vawebwork"
)

const (
	appName = "MY-APP"
)

func main() {
	// defining default values
	var (
		logsPath  string    = vafswork.GetExePath() + "/logs" + "_" + appName
		startTime time.Time = time.Now()
		// mailingFile       string = getExePath() + "/data/mailing.json"
	)

	// flags
	logsDir := flag.String("log-dir", logsPath, "set custom log dir")
	logsToKeep := flag.Int("keep-logs", 7, "set number of logs to keep after rotation")

	flag.Usage = func() {
		fmt.Println("THIS APP IS FOR ...")
		fmt.Println("Usage: <app> [-opt] ...")
		fmt.Println("Flags:")
		flag.PrintDefaults()
	}

	flag.Parse()

	// logging
	logFile, err := logging.StartLogging(appName, *logsDir, *logsToKeep)
	if err != nil {
		log.Fatalf("failed to start logging:\n\t%s", err)
	}
	defer logFile.Close()

	// starting programm notification
	log.Printf("%s Started", appName)

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
