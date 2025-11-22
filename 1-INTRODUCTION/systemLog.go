// UNIX
package main

import (
	"log"
	"log/syslog"
)

func systemLog() {
	// Create a new connection to the system logger
	sysLog, err := syslog.New(syslog.LOG_INFO|syslog.LOG_LOCAL7, "myGoApp")
	if err != nil {
		log.Println("Error creating syslog:", err)
		return
	}

	// Send all log output to syslog
	log.SetOutput(sysLog)

	// Write a log message
	log.Print("Everything is fine!")
}

// SYSLOG IS NOT SUPPORTTED IN WINDOWS
