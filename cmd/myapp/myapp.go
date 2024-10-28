package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"time"

	// change this path for your project
	"github.com/slayerjk/go-project-template/internal/logging"
	"github.com/slayerjk/go-project-template/internal/rotatefiles"
	// "github.com/slayerjk/go-project-template/internal/mailing"
)

// log default path & logs to keep after rotation
const (
	appName = "MYAPP"
)

func main() {
	// get executable's working dir
	exe, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	exePath := filepath.Dir(exe)

	// defining default values
	var (
		startTime         time.Time = time.Now()
		defaultLogPath    string    = exePath + "/logs"
		defaultLogsToKeep int       = 3
		// mailingFile       string = exePath + "/data/mailing.json"
	)

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
	log.Printf("%s Started", appName)

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
