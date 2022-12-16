package firewall

import (
	"bytes"
	"context"
	"errors"
	"go2ban/pkg/config"
	"log"
	"os/exec"
	"time"
)

func iptablesBlock(ctx context.Context, ip string) {
	err := runCMD("iptables --table raw --append go2ban --source " + ip + " --jump DROP")
	if err != nil {
		log.Println("Not blocked ", ip, err)
	}
}

func workerIptables() {
	err := runCMD("iptables --table raw --new go2ban")
	if err != nil && err.Error() != "exit status 1" {
		log.Println("Not add chain go2ban ", err)
	}
	byt, err := runOutputCMD("iptables-save")
	if len(byt) == 0 {
		log.Fatalln("Can't get iptables settings, iptables-save", err)
	}
	if !bytes.Contains(byt, []byte{'j', ' ', 'g', 'o', '2', 'b', 'a', 'n'}) {
		err = runCMD("iptables --table raw --insert PREROUTING --jump go2ban")
		if err != nil {
			log.Println("Not add chain go2ban to table raw ", err)
		}
	}
	log.Println("Iptables: add chain go2ban to table raw")
	go func() {
		for {
			count := countBlocked()
			cfgMaxLocked := config.Get().BlockedIps
			if count > 0 && cfgMaxLocked < count {
				start := time.Now()
				for i := 0; i < (count-cfgMaxLocked)+cfgMaxLocked/10; i++ {
					err = runCMD("iptables --table raw --delete go2ban 1")
					if err != nil && err.Error() != "exit status 1" {
						log.Println("Can't del ip ", err)
					}
				}
				log.Printf("Worker Iptables clear %d in %.2f seconds", count, time.Since(start).Seconds())
			}
			time.Sleep(time.Hour * sleepHour)
		}
	}()
}

func iptablesUnlockAll(ctx context.Context) (ips int, err error) {
	ips = countBlocked()
	err = runCMD(`iptables --table raw --flush go2ban`)
	if err != nil && err.Error() != "exit status 1" {
		log.Println("Don't del list ", err)
		return 0, errors.New("Not found chain go2ban")
	}
	log.Println("Iptables: clear all")
	return
}

func countBlocked() (ips int) {
	byt, err := runOutputCMD("iptables-save")
	if err == nil {
		ips = bytes.Count(byt, []byte{'A', ' ', 'g', 'o', '2', 'b', 'a', 'n'})
	}
	return
}

func runCMD(firewallCMD string) error {
	err := exec.Command("sh", "-c", firewallCMD).Run()
	return err
}

func runOutputCMD(firewallCMD string) ([]byte, error) {
	b, err := exec.Command("sh", "-c", firewallCMD).Output()
	return b, err
}
