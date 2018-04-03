package irsync

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"time"
)

// Logger provides an interface for any logging
// packing providing a Info(args ...interface{}) method
type Logger interface {
	Info(args ...interface{})
}

// Status provides access to the current state of
// sync and interval status.
type Status struct {
	CurrentInterval int
	LineN           int
}

// Sync configuration
type Sync struct {
	Log             Logger
	ActivityTimeout time.Duration
	Interval        time.Duration
	Status          Status
	Flags           string
	Delete          bool   // rsync --delete
	LocationFrom    string // rsync location from
	LocationTo      string // srync location to
}

// IntervalRSync endlessly runs RSync on interval
func (s *Sync) IntervalRSync() {
	// ensure we have an initialized status struct
	s.initStatus()

	s.Log.Info("Interval every %d seconds.", s.Interval/time.Second)
	s.Log.Info("Starting interval %d.", s.Status.CurrentInterval)
	s.Log.Info("Starting interval with timeout set for %d seconds with no activity.", s.ActivityTimeout/time.Second)

	// run rsync
	s.RSync()

	s.Log.Info("Interval complete. Waiting %d seconds.", s.Interval/time.Second)
	time.Sleep(s.Interval)

	s.Status.CurrentInterval++
	// interval
	s.IntervalRSync()
}

// RSync runs the command rsync
func (s *Sync) RSync() {
	// ensure we have an initialized status struct
	s.initStatus()
	s.Status.LineN = 0

	// initialize command args
	args := s.getArgs()

	// Create Cmd with options
	cmd := exec.Command("rsync", args...)

	// output command we are using
	s.Log.Info("rsync args %s", cmd.Args)

	line := make(chan string)
	done := make(chan bool)
	go s.runCommand(cmd, line, done)

	for {
		select {
		case l := <-line:
			s.Status.LineN++
			info := make(map[string]interface{})
			info["line_n"] = s.Status.LineN
			info["interval"] = s.Status.CurrentInterval
			info["line"] = l
			s.Log.Info(info)

		case <-done:
			s.Log.Info("Done with RSync.")
			return

		case <-time.After(s.ActivityTimeout): // no file should take loner than this
			s.Log.Info("rsync activity timeout.")
			return
		}
	}
}

// getArgs get rsync args
func (s *Sync) getArgs() []string {
	args := make([]string, 0)

	// rsync flags
	args = append(args, s.Flags)

	// rsync --delete
	if s.Delete {
		args = append(args, "--delete")
	}

	// rsync from
	args = append(args, s.LocationFrom)

	// rsync to
	args = append(args, s.LocationTo)

	return args
}

// initStatus initialize status if empty
func (s *Sync) initStatus() {
	if s.Status == (Status{}) {
		s.Status = Status{
			CurrentInterval: 1,
			LineN:           0,
		}
	}
}

// runCommand is a generic command runner
func (s *Sync) runCommand(cmd *exec.Cmd, line chan string, done chan bool) {

	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating StdoutPipe for Cmd", err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(cmdReader)
	go func() {
		for scanner.Scan() {
			line <- scanner.Text()
		}
	}()

	err = cmd.Start()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error starting Cmd", err)
		os.Exit(1)
	}

	err = cmd.Wait()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error waiting for Cmd", err)
		os.Exit(1)
	}

	done <- true
}
