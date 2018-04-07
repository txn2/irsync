package main

import "testing"

// TestCmdParsing test command line parsing.
func TestCmdParsing(t *testing.T) {
	commandSets := []struct {
		Args          []string
		RsyncExpects  int
		IrsyncExpects int
	}{
		{
			Args: []string{
				"-pvrt",
				"--modify-window=30",
				"--delete",
				"--exclude='custom'",
				"--exclude='somefile.txt'",
				"rsync://test@mk.imti.co:873/data/",
				"/media",
			},
			RsyncExpects:  7,
			IrsyncExpects: 0,
		},
		{
			Args: []string{
				"-pvrt",
				"--modify-window=30",
				"--irsync-timeout-seconds=200",
				"--irsync-interval-seconds=300",
				"--delete",
				"--exclude='custom'",
				"--exclude='somefile.txt'",
				"rsync://test@mk.imti.co:873/data/",
				"/media",
			},
			RsyncExpects:  7,
			IrsyncExpects: 2,
		},
	}

	for _, commandSet := range commandSets {

		rsyncArgs, irsyncArgs := parseCmdArgs(commandSet.Args)

		// Test sorting out rsync from irsync flags
		//
		if len(rsyncArgs) != commandSet.RsyncExpects {
			t.Errorf("rsync got %d args but expected %d.", len(rsyncArgs), commandSet.RsyncExpects)
		}

		if len(irsyncArgs) != commandSet.IrsyncExpects {
			t.Errorf("irsync got %d args but expected %d.", len(irsyncArgs), commandSet.IrsyncExpects)
		}

		// Test default and specified values
		if getIntegerFromArgs("--test-non-exist-key", irsyncArgs, 999) != 999 {
			t.Errorf("Got wrong value or default value from args.")
		}

		kvt := getIntegerFromArgs("--irsync-timeout-seconds", irsyncArgs, 888)
		if !(kvt == 888 || kvt == 200) {
			t.Errorf("Got wrong value or default value from args, got %d.", kvt)
		}

		kvi := getIntegerFromArgs("--irsync-interval-seconds", irsyncArgs, 887)
		if !(kvi == 887 || kvi == 300) {
			t.Errorf("Got wrong value or default value from args, got %d.", kvi)
		}

	}
}
