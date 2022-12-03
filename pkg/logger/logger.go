package logger

import (
	"go2ban/pkg/config"
	"go2ban/pkg/osUtil"
	"log"
	"log/syslog"
	"os"
	"path/filepath"
)

const logExp = ".log"

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

var logFile *os.File

func Start() {
	logFilePath := filepath.Join(config.Get().LogDir, osUtil.DateNow()+logExp)
	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalln("Dont create log file")
	}
	log.SetOutput(logFile)
}

func SendSyslogMail(msg string) {
	syslog, err := syslog.New(syslog.LOG_MAIL, "go2ban")
	if err != nil {
		log.Println("Load syslog ", err)
	} else {
		log.SetOutput(syslog)
		log.Println(msg)
		log.SetOutput(logFile)
	}
}
