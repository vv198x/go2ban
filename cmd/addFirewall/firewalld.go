package addFirewall

import (
	"fmt"
	"go2ban/pkg/logger"
	"log"
	"os/exec"
)

func firewalldBlock(ip string) {
	firewallCMD := fmt.Sprintf(
		`firewall-cmd --permanent --add-rich-rule="rule family='ipv4' source address='%v' drop"`, ip)
	b, err := exec.Command("sudo", "bash", "-c", firewallCMD).Output()
	if err != nil {
		log.Fatalln(err)
	}
	if string(b) == "success\n" {
		firewallCMD = "firewall-cmd --reload"
		err = exec.Command("sudo", "bash", "-c", firewallCMD).Run()
		if err != nil {
			log.Println(err)
		}
		logger.SendSyslogMail("BANED: " + ip)
	} else {
		log.Println("Dont add address to rule ", ip)
	}
}
