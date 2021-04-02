package utils

import (
	"log"
	"os"
)

var InfoLog *log.Logger
var DebugLog *log.Logger
var ErrorLog *log.Logger

func init() {
	InfoLog = &log.Logger{}
	InfoLog.SetPrefix("[INFO]  ")
	InfoLog.SetOutput(os.Stdout)

	DebugLog = &log.Logger{}
	DebugLog.SetPrefix("[DEBUG] ")
	DebugLog.SetOutput(os.Stdout)

	ErrorLog = &log.Logger{}
	ErrorLog.SetPrefix("[ERROR] ")
	ErrorLog.SetOutput(os.Stderr)
}
