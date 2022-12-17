package REST

import (
	"go2ban/pkg/config"
	"log"
	"net/http"
	"strconv"
)

func Start(runAsDaemon bool) {
	if runAsDaemon {
		go func() {
			port := config.Get().RestPort
			if port == "" {
				return
			}

			_, err := strconv.Atoi(port)
			if err != nil {
				log.Println("Wrong port REST Server ", err)
				return
			}

			err = http.ListenAndServe(":"+port, nil)
			if err != nil {
				log.Fatalln("Can't start REST server ", err)
			}
		}()
	}
}
