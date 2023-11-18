package main

import "os"

const (
	// exit is exit code which is returned by realMain function.
	// exit code is passed to os.Exit function.
	exitOK int = iota
	exitError
)

func main() {
	os.Exit(realMain(os.Args))
}

func realMain(args []string) int {

	return exitOK
}
