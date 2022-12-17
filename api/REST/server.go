package REST

import (
	"go2ban/pkg/config"
	"log"
	"net/http"
)

func Start(runAsDaemon bool) {
	if runAsDaemon {
		go func() {
			//TODO check port
			err := http.ListenAndServe(":"+config.Get().RestPort, nil)

			if err != nil {
				log.Fatalln("Can't start REST server")
			}
		}()
	}
}
