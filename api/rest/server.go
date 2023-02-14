package rest

import (
	"github.com/vv198x/go2ban/config"
	"log"
	"net/http"
	"regexp"
)

func Start(runAsDaemon bool) {
	if runAsDaemon {
		go func() {
			port := config.Get().RestPort
			if port == "" {
				return
			}

			if !regexp.MustCompile(`(\d)+`).MatchString(port) {
				log.Println("Wrong port REST Server ")
				return
			}

			err := http.ListenAndServe(":"+port, nil)
			if err != nil {
				log.Fatalln("Can't start REST server ", err)
			}
		}()
	}
}
