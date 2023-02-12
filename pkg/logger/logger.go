package logger

import (
	"github.com/vv198x/go2ban/config"
	"log"
	"os"
	"path/filepath"
	"time"
)

const logExp = ".log"

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func Start() {
	var logFile *os.File
	logFilePath := filepath.Join(
		config.Get().LogDir,
		time.Now().Format("06.01.02")+logExp)
	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalln("Dont create log file")
	}
	log.SetOutput(logFile)
}
