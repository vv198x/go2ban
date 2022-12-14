package REST

import (
	"context"
	"encoding/json"
	"go2ban/cmd/firewall"
	"go2ban/pkg/validator"
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
		json.NewEncoder(writer).Encode(sayOk{Message: "Bad Request"})
		log.Println("REST Bad Request", request.Method, request.URL)
		return
	}

	ip, err := validator.CheckIp(ip)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(sayOk{Message: err.Error()})
		log.Println("REST validator:", err)
	} else {

		go firewall.BlockIP(context.Background(), ip)

		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode(sayOk{true, "Ip blocked " + ip})
		log.Println("REST ip blocked:", ip)
	}
}
