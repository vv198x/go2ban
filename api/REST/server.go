package REST

import (
	"context"
	"go2ban/cmd/firewall"
	"go2ban/pkg/config"
	"go2ban/pkg/validator"
	"log"
	"net/http"
)

func Start(runAsDaemon bool) {
	if runAsDaemon {
		http.HandleFunc("/", getIp)

		go func() {
			err := http.ListenAndServe(":"+config.Get().RestPort, nil)
			if err != nil {
				log.Fatalln("Can't start REST server")
			}
		}()

	}
}
func getIp(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodPost {
		ip, err := validator.CheckIp(request.PostFormValue("ip"))
		if err == nil {
			ctx := context.Background()
			go firewall.BlockIP(ctx, ip)
			writer.WriteHeader(http.StatusOK)
			writer.Header().Set("Content-Type", "application/json")
			log.Println("REST ip blocked:", ip)
		} else {
			writer.WriteHeader(http.StatusBadRequest)
			log.Println("REST validator:", err)
		}
	}
	writer.WriteHeader(http.StatusBadRequest)
}
