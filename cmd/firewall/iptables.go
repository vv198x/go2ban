package firewall

import (
	"bytes"
	"errors"
	"log"
	"os/exec"
)

func iptablesBlock(ip string) {
	err := runCMD("iptables -A go2ban -s " + ip + " -j DROP")
	if err != nil {
		log.Println("Not blocked ", ip, err)
	}
}

func workerIptables() {
	err := runCMD("iptables -N go2ban")
	if err != nil && err.Error() != "exit status 1" {
		log.Println("Not add chain go2ban ", err)
	}
	byt, err := runOutputCMD("iptables-save")
	if len(byt) == 0 {
		log.Fatalln("Can't get iptables settings, iptables-save", err)
	}
	if !bytes.Contains(byt, []byte{'j', ' ', 'g', 'o', '2', 'b', 'a', 'n'}) {
		err := runCMD("iptables -A INPUT -j go2ban")
		if err != nil {
			log.Println("Not add chain go2ban to table input ", err)
		}
	}
}

func firewalldUnlockAll() (ips int, err error) {
	byt, err := runOutputCMD("iptables-save")
	if err == nil {
		ips = bytes.Count(byt, []byte{'A', ' ', 'g', 'o', '2', 'b', 'a', 'n'})
	}
	err = runCMD(`iptables -F go2ban`)
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
