package logger

import (
	"log"
	"log/syslog"
	"microservice2ban/pkg/config"
	"microservice2ban/pkg/osUtil"
	"os"
	"path/filepath"
	"reflect"
)

const logExp = ".log"

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

var logFile *os.File

func Start() {
	logFilePath := filepath.Join(config.Get("logDir"), osUtil.DateNow()+logExp)
	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalln("Dont create log file")
	}
	log.SetOutput(logFile)
}

type Empty struct{}

func SendSyslogMail(msg string) {
	syslog, err := syslog.New(syslog.LOG_MAIL, reflect.TypeOf(Empty{}).PkgPath())
	if err != nil {
		log.Println("Load syslog ", err)
	} else {
		log.SetOutput(syslog)
		log.Println(msg)
		log.SetOutput(logFile)
	}
}
