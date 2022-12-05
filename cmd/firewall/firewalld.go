package firewall

import (
	"errors"
	"fmt"
	"go2ban/pkg/logger"
	"log"
	"os/exec"
	"regexp"
)

func firewallBlock(ip string) {
	firewallCMD := fmt.Sprintf(
		`firewall-cmd --permanent --add-rich-rule="rule family='ipv4' source address='%v' drop"`, ip)
	byteArr, err := runOutputFirewalld(firewallCMD)
	if err == nil && string(byteArr) == "success\n" {
		err = reloadFirewalld()
		logger.SendSyslogMail("BANED: " + ip)
	} else {
		log.Println("Dont add address to rule ", ip, err)
	}
}

func firewalldUnlockAll() (ips int, err error) {
	byteArr, err := runOutputFirewalld(`firewall-cmd  --list-rich-rules`)
	if err != nil {
		log.Println("firewalldUnlockAll get list ", err)
		return
	}
	IPst := regexp.MustCompile(`((25[0-5]|(2[0-4]|1\d|[1-9]|)\d)\.?\b){4}`).FindAll(byteArr, -1)
	countIP := len(IPst)
	if countIP == 0 {
		return len(IPst), errors.New("Rules not found")
	} else {
		for _, ip := range IPst {
			firewallCMD := fmt.Sprintf(
				`firewall-cmd --permanent --remove-rich-rule="rule family='ipv4' source address='%s' drop"`, ip)
			byteArr, err = runOutputFirewalld(firewallCMD)
			if err != nil && string(byteArr) != "success\n" {
				log.Println("Dont remove-rich-rule ", string(byteArr))
				return len(IPst), errors.New("Can't delete rule with ip address: " + string(ip))
			}
			countIP--
		}
		if countIP == 0 {
			err = reloadFirewalld()
			if err == nil {
				return len(IPst), nil
			}
		}
	}
	return countIP, err
}

func reloadFirewalld() error {
	firewallCMD := "firewall-cmd --reload"
	err := runCmdFirewalld(firewallCMD)
	if err != nil {
		log.Println("Dont reload firewalld ", err)
	}
	return err
}

func runCmdFirewalld(firewallCMD string) error {
	err := exec.Command("sudo", "bash", "-c", firewallCMD).Run()
	if err != nil {
		log.Println("runCmdFirewalld", err)
	}
	return err
}
func runOutputFirewalld(firewallCMD string) ([]byte, error) {
	b, err := exec.Command("sudo", "bash", "-c", firewallCMD).Output()
	if err != nil {
		log.Println("runOutputFirewalld ", err)
	}
	return b, err
}
