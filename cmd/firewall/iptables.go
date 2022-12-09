package firewall

import (
	"bytes"
	"context"
	"errors"
	"log"
	"os/exec"
)

func iptablesBlock(ctx context.Context, ip string) {
	err := runCMD("iptables --table raw --append go2ban --source " + ip + " --jump DROP")
	if err != nil {
		log.Println("Not blocked ", ip, err)
	}
}

func workerIptables() { //TODO add clear list
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
}

func firewalldUnlockAll() (ips int, err error) {
	byt, err := runOutputCMD("iptables-save")
	if err == nil {
		ips = bytes.Count(byt, []byte{'A', ' ', 'g', 'o', '2', 'b', 'a', 'n'})
	}
	err = runCMD(`iptables --table raw --flush go2ban`)
	if err != nil && err.Error() != "exit status 1" {
		log.Println("Don't del list ", err)
		return 0, errors.New("Not found chain go2ban")
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
