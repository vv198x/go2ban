package logger

import (
	"log"
	"microservice2ban/pkg/config"
	"microservice2ban/pkg/osUtil"
	"os"
	"path/filepath"
)

const logExp = ".log"

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func Start() {
	logFilePath := filepath.Join(
		config.Get("logDir"),
		osUtil.DateNow()+logExp)
	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalln("Dont create log file")
	}
	log.SetOutput(logFile)

}
