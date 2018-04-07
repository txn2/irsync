package main

import (
	"os"
	"strings"

	"strconv"

	"fmt"
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

	syncLog, err := bunyan.CreateLogger(logConfig)
	if err != nil {
		panic(err)
	}

	// rsync has over 100 arguments, lets add a couple more
	rsyncArgs, irsyncArgs := parseCmdArgs(os.Args[1:])

	// default activity timeout
	activityTimeout := 7200
	if secs, ok := irsyncArgs["--irsync-timeout-seconds"]; ok {
		at, err := strconv.Atoi(secs)
		if err != nil {
			fmt.Println("value for --irsync-timeout-seconds must be a postive integer (seconds).")
			os.Exit(1)
		}

		activityTimeout = at
	}

	// default interval
	interval := 120
	if secs, ok := irsyncArgs["--irsync-interval-seconds"]; ok {
		itvl, err := strconv.Atoi(secs)
		if err != nil && itvl > 0 {
			fmt.Println("value for --irsync-interval-seconds must be a postive integer (seconds).")
			os.Exit(1)
		}

		interval = itvl
	}

	fmt.Printf("Timeout: %d, Interval: %d\n", activityTimeout, interval)

	sync := irsync.Sync{
		Log:             &syncLog,
		ActivityTimeout: time.Duration(activityTimeout) * time.Second,
		Interval:        time.Duration(interval) * time.Second,
		RsyncArgs:       rsyncArgs,
	}

	// run rsync on interval (block)
	<-sync.Run()

}

// parseArgs pulls out irsync specific args
func parseCmdArgs(args []string) ([]string, map[string]string) {
	rsyncArgs := []string{}
	irsyncArgs := map[string]string{}

	for _, arg := range args {
		if strings.HasPrefix(arg, "--irsync-") {
			if a := strings.Split(arg, "="); len(a) > 1 {
				irsyncArgs[a[0]] = a[1]
			}
			continue
		}

		rsyncArgs = append(rsyncArgs, arg)
	}

	return rsyncArgs, irsyncArgs
}
