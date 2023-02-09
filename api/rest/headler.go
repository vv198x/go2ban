package rest

import (
	"context"
	"encoding/json"
	"github.com/vv198x/go2ban/cmd/firewall"
	"github.com/vv198x/go2ban/pkg/validator"
	"log"
	"net/http"
)

type sayOk struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
}

func init() {
	http.HandleFunc("/", getIp)
}

func getIp(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	ip := request.PostFormValue("ip")

	if request.Method != http.MethodPost || ip == "" {
		writer.WriteHeader(http.StatusBadRequest)
		if errJ := json.NewEncoder(writer).Encode(sayOk{Message: "Bad Request"}); errJ != nil {
			log.Println(errJ)
		}
		log.Println("REST Bad Request", request.Method, request.URL)
		return
	}

	ip, err := validator.CheckIp(ip)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		if errJ := json.NewEncoder(writer).Encode(sayOk{Message: err.Error()}); errJ != nil {
			log.Println(errJ)
		}
		log.Println("REST validator: ", err)
		return
	}

	go firewall.BlockIP(context.Background(), ip)

	writer.WriteHeader(http.StatusOK)
	if errJ := json.NewEncoder(writer).Encode(sayOk{true, "Ip blocked " + ip}); errJ != nil {
		log.Println(errJ)
	}
	log.Println("REST ip blocked:", ip)

}
