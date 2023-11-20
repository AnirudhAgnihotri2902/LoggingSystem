package main

import (
	logstash_logger "github.com/AnirudhAgnihotri2902/logging-system/app/logger"
	"os"
	"time"
)

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
	logger := logstash_logger.Init("localhost", 9600, "tcp", 5)

	payload := map[string]interface{}{
		"message":     "TEST_MSG",
		"error":       false,
		"code":        200,
		"resource_id": "test_resource_id",
		"timestamp":   time.Now(),
		"trace_id":    "test_trace_id",
		"span_id":     "test_span_id",
		"commit":      "test_commit",
		"meta_data":   "meta_data",
	}

	logger.Log(payload)   // Generic log
	logger.Info(payload)  // Adds "severity": "INFO"
	logger.Debug(payload) // Adds "severity": "DEBUG"
	logger.Warn(payload)  // Adds "severity": "WARN"
	logger.Error(payload) // Adds "severity": "ERROR"
	return exitOK
}
