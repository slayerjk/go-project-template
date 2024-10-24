package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	// change this path for your project
	"template/internal/logging"
	"template/internal/mailing"
	"template/internal/rotatefiles"
)

// log default path & logs to keep after rotation
const (
	defaultLogPath    = "logs"
	defaultLogsToKeep = 3
)

func main() {
	// flags
	logDir := flag.String("log-dir", defaultLogPath, "set custom log dir")
	logsToKeep := flag.Int("keep-logs", defaultLogsToKeep, "set number of logs to keep after rotation")
	flag.Parse()

	// logging
	appName := "MYAPP"

	logFile, err := logging.StartLogging(appName, *logDir, 3)
	if err != nil {
		log.Fatalf("failed to start logging:\n\t%s", err)
	}
	defer logFile.Close()

	// starting programm notification
	startTime := time.Now()
	log.Println("Program Started")

	// main code here

	/* // mailing example 'report'(read and send log file)
	report, err := os.ReadFile(logFile.Name())
	if err != nil {
		log.Fatal(err)
	}
	errM1 := mailing.SendPlainEmailWoAuth("mailing.json", "report", appName, report, startTime)
	if errM1 != nil {
		log.Printf("failed to send email:\n\t%v", errM1)
	} */

	// mailing example 'error'(just error text)
	newError := fmt.Errorf("custom error")
	errM2 := mailing.SendPlainEmailWoAuth("mailing.json", "error", appName, []byte(newError.Error()), startTime)
	if errM2 != nil {
		log.Printf("failed to send email:\n\t%v", errM2)
	}

	// count & print estimated time
	endTime := time.Now()
	log.Printf("Program Done\n\tEstimated time is %f seconds", endTime.Sub(startTime).Seconds())

	// close logfile and rotate logs
	logFile.Close()

	if err := rotatefiles.RotateFilesByMtime(*logDir, *logsToKeep); err != nil {
		log.Fatalf("failed to rotate logs:\n\t%s", err)
	}
}
