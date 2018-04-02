package main

import (
	"os"
	"time"

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

	sync := irsync.Sync{
		Log:             &blog,
		ActivityTimeout: 2 * time.Hour,    // no file should take more than 2 hours
		Interval:        10 * time.Second, // run again in 10 seconds after rsync completes
		LocationFrom:    getEnv("IRSYNC_FROM", "rsync://byp@sync.byp.mobi:31873/data/"),
		LocationTo:      getEnv("IRSYNC_TO", "./data"),
		Flags:           getEnv("IRSYNC_FLAGS", "-avzr"),
		Delete:          false,
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
