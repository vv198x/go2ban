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
			safeBlockBlackListIPs(apiKey, urlBlacklist)
			<-ticker.C
		}
	}()

}

// safeBlockBlackListIPs isolates the ticked goroutine from a panic in
// blockBlackListIPs (e.g. a transient abuseipdb outage) so it can't take
// down the whole daemon.
func safeBlockBlackListIPs(apiKey, urlBl string) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("blockBlackListIPs recovered from panic:", r)
		}
	}()
	blockBlackListIPs(apiKey, urlBl)
}

func blockBlackListIPs(apiKey string, urlBl string) {
	// Number of results to return (free max 10000)
	// Approximately 33% match in 2000 ips
	limit := config.Get().AbuseipdbIPs

	// Minimum abuse confidence score to return (0-100)
	minimumScore := 90

	// Send GET request
	req, err := http.NewRequest(http.MethodGet,
		fmt.Sprintf("%s?confidenceMinimum=%d&limit=%d", urlBl, minimumScore, limit), nil)
	if err != nil {
		log.Println("Build req error", err)
	}
	req.Header.Set("Key", apiKey)
	req.Header.Set("Accept", "text/plain")

	client := &http.Client{Timeout: time.Second * 10}
	var resp *http.Response

	// Three attempts
	for i := 0; i <= 3; i++ {
		resp, err = client.Do(req)
		if err != nil {
			log.Printf("Send req Do error: %v, retrying %d...", err, i)
			time.Sleep(time.Minute)
			continue
		}
		break
	}

	// All attempts failed (network/DNS/TLS error) - resp is nil here.
	if err != nil || resp == nil {
		log.Println("Send req Do error, giving up:", err)
		return
	}
	defer resp.Body.Close()

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
