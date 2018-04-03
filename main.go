package main

import (
	"os"
	"time"

	"strconv"

	"github.com/bhoriuchi/go-bunyan/bunyan"
	"github.com/cjimti/irsync/irsync"
)

func main() {

	logConfig := bunyan.Config{
		Name:   "irsync",
		Stream: os.Stdout,
		Level:  bunyan.LogLevelDebug,
	}

	blog, err := bunyan.CreateLogger(logConfig)
	if err != nil {
		panic(err)
	}

	blog.Info("Starting irsync...")

	envIntervalSeconds := getEnv("IRSYNC_INTERVAL", "10") // default 10 seconds
	intervalSeconds, err := strconv.Atoi(envIntervalSeconds)
	if err != nil {
		blog.Info("IRSYNC_INTERVAL can not be converted to a number representing seconds.")
	}

	envTimeoutSeconds := getEnv("IRSYNC_TIMEOUT", "7200") // default 2 hours
	timeoutSeconds, err := strconv.Atoi(envTimeoutSeconds)
	if err != nil {
		blog.Info("IRSYNC_TIMEOUT can not be converted to a number representing seconds.")
	}

	deleteFiles := false

	encDelete := getEnv("IRSYNC_DELETE", "false")
	if encDelete == "true" {
		deleteFiles = true
	}

	sync := irsync.Sync{
		Log:             &blog,
		ActivityTimeout: time.Duration(timeoutSeconds) * time.Second,
		Interval:        time.Duration(intervalSeconds) * time.Second,
		LocationFrom:    getEnv("IRSYNC_FROM", "./"),
		LocationTo:      getEnv("IRSYNC_TO", "./data"),
		Flags:           getEnv("IRSYNC_FLAGS", "-ogplvrt"),
		Delete:          deleteFiles,
	}

	// run rsync on interval (blocking)
	sync.IntervalRSync()

	// should never get here if interval is used
	blog.Info("exit.")
}

// getEnv gets an environment variable or sets a default if
// one does not exist.
func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}

	return value
}
