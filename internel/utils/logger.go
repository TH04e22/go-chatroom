package utils

import (
	"log"
	"os"
)

var Info_log *log.Logger
var Debug_log *log.Logger
var Error_log *log.Logger

func init() {
	// a basic logger
	Info_log = &log.Logger{}
	Info_log.SetPrefix("[INFO]  ")
	Info_log.SetOutput(os.Stdout)

	Debug_log = &log.Logger{}
	Debug_log.SetPrefix("[DEBUG] ")
	Debug_log.SetOutput(os.Stdout)

	Error_log = &log.Logger{}
	Error_log.SetPrefix("[ERROR] ")
	Error_log.SetOutput(os.Stderr)
}
