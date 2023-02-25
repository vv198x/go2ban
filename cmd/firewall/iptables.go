package firewall

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/vv198x/go2ban/config"
	"log"
	"strings"
	"time"
)

type iptables struct{}

func (fw *iptables) Block(ctx context.Context, ip string) {
	start := time.Now()

	err := runCMD("iptables --table raw --append go2ban --source " + ip + " --jump DROP")
	if err != nil {
		log.Println("Not blocked ", ip, err)
	}

	log.Println("Blocked in milliseconds: ", time.Since(start).Milliseconds())
}

func (fw *iptables) Worker() {

	err := runCMD("iptables --table raw --new go2ban")
	if err != nil && err.Error() != "exit status 1" {
		log.Println("Not add chain go2ban ", err)
	}

	byt, err := runOutputCMD("iptables-save")
	if len(byt) == 0 {
		log.Fatalln("Can't get iptables settings, iptables-save", err)
	}
	if !bytes.Contains(byt, []byte("j go2ban")) {
		err = runCMD("iptables --table raw --insert PREROUTING --jump go2ban")
		if err != nil {
			log.Println("Not add chain go2ban to table raw ", err)
		}
	}
	log.Println("Iptables: add chain go2ban to table raw")
	go func() {
		for {
			count := fw.countBlocked()
			cfgMaxLocked := config.Get().BlockedIps
			if count > 0 && cfgMaxLocked < count {
				start := time.Now()

				// Delete ip up to (maximum allowed - 10%). One at a time, old ones first
				beRemoved := (count - cfgMaxLocked) + cfgMaxLocked/10
				for i := 0; i < beRemoved; i++ {
					err = runCMD("iptables --table raw --delete go2ban 1")
					if err != nil && err.Error() != "exit status 1" {
						log.Println("Can't del ip ", err)
					}
				}
				log.Printf("Worker Iptables clear %d in %.2f seconds", beRemoved, time.Since(start).Seconds())
			}
			time.Sleep(time.Hour * config.WorkerSleepHour)
		}
	}()
}

func (fw *iptables) UnlockAll(ctx context.Context) (ips int, err error) {
	ips = fw.countBlocked()
	err = runCMD(`iptables --table raw --flush go2ban`)
	if err != nil && err.Error() != "exit status 1" {
		log.Println("Don't del list ", err)
		return 0, errors.New("not found chain go2ban")
	}
	log.Println("Iptables: clear all")
	return
}

func (fw *iptables) countBlocked() (ips int) {
	byt, err := runOutputCMD("iptables-save")
	if err == nil {
		ips = bytes.Count(byt, []byte("A go2ban"))
	}
	return
}

func (fw *iptables) GetBlocked() map[string]struct{} {
	byt, err := runOutputCMD("iptables-save")
	if err != nil {
		log.Println("iptables-save error", err)
		return nil
	}
	m := make(map[string]struct{})

	for _, st := range bytes.Split(byt, []byte("\n")) {
		var buf string
		fmt.Sscanf(string(st), "-A go2ban -s %s/32 -j DROP", &buf)
		m[strings.TrimSuffix(buf, "/32")] = struct{}{}
	}
	return m
}
