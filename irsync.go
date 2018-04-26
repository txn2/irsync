package main

import (
	"os"
	"strings"

	"strconv"

	"fmt"
	"time"

	"github.com/bhoriuchi/go-bunyan/bunyan"
	"github.com/txn2/irsync/irsync"
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
	activityTimeout := getIntegerFromArgs("--irsync-timeout-seconds", irsyncArgs, 7200)

	// default interval
	interval := getIntegerFromArgs("--irsync-interval-seconds", irsyncArgs, 120)

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

// getIntegerFromArgs returns an integer if key exists or falls back
// os exits if key exists and cannot be converted to an int
func getIntegerFromArgs(key string, args map[string]string, fallback int) int {
	v := fallback

	if nvStr, ok := args[key]; ok {
		nv, err := strconv.Atoi(nvStr)
		if err != nil {
			fmt.Printf("value for %s must be a postive integer got %d.", key, nv)
			os.Exit(1)
		}

		v = nv
	}

	return v
}

// parseArgs pulls out irsync specific args
func parseCmdArgs(args []string) ([]string, map[string]string) {
	rsyncArgs := make([]string, 0)
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
