package abuseipdb

import (
	"bytes"
	"context"
	"fmt"
	"github.com/vv198x/go2ban/cmd/firewall"
	"github.com/vv198x/go2ban/cmd/validator"
	"github.com/vv198x/go2ban/config"
	"io"
	"log"
	"net/http"
	"regexp"
	"time"
)

const urlBlacklist = "https://api.abuseipdb.com/api/v2/blacklist/"

func Scheduler(apiKey string) {
	if !config.Get().Flags.RunAsDaemon || !regexp.MustCompile(`[\d\w]{80}`).MatchString(apiKey) {
		return
	}
	go func() {
		ticker := time.NewTicker(config.WorkerSleepHour * time.Hour)
		for {
			go blockBlackListIPs(apiKey, urlBlacklist)
			<-ticker.C
		}
	}()

}

func blockBlackListIPs(apiKey string, urlBl string) {
	// Number of results to return (free max 10000)
	// Approximately 33% match in 2000 ips
	limit := config.Get().AbuseipdbIPs

	// Minimum abuse confidence score to return (0-100)
	minimumScore := 90

	// Send GET request
	ctx, stop := context.WithTimeout(context.Background(), time.Second*10)
	defer stop()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet,
		fmt.Sprintf("%s?confidenceMinimum=%d&limit=%d", urlBl, minimumScore, limit), nil)
	if err != nil {
		log.Println("Build req error", err)
	}
	req.Header.Set("Key", apiKey)
	req.Header.Set("Accept", "text/plain")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Send req Do error", err)
	}
	if resp.Body != nil {
		defer resp.Body.Close()
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Read body error", err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Println("Abuseipdb say error", string(body))
		return
	}

	start := time.Now()
	// map already blocked, so as not to block 2 times
	m := firewall.Do().GetBlocked()
	var C int
	sts := bytes.Split(body, []byte("\n"))
	for _, st := range sts {
		ip := string(st)
		if _, find := m[ip]; !find && ip != "" {
			ip, err = validator.CheckIp(ip)
			if err != nil {
				log.Println("abuseipdb validator err", err)
				continue
			}
			firewall.Do().Block(context.Background(), ip)
			C++
		}
	}
	log.Printf("End abuseipdb second:%.2f, get IPs:%d, new IPs: %d", time.Since(start).Seconds(), len(sts)-1, C)
}
