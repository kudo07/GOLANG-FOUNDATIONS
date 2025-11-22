package main

import (
	"log"
	"log/syslog"
)

func ststemLog() {
	sysLog, err := syslog.New(syslog.LOG_INFO|syslog.LOG_LOCAL7, "myGoApp")
	if err != nil {
		log.Println("Error creating syslog:", err)
		return
	}
}
