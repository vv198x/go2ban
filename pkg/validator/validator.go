package validator

import (
	"errors"
	"microservice2ban/pkg/osUtil"
	"regexp"
)

func CheckIp(target string) error {
	localAddress := osUtil.GetLocalIPs()
	if regexp.MustCompile(`^((25[0-5]|(2[0-4]|1\d|[1-9]|)\d)\.?\b){4}$`).MatchString(target) {
		for _, addr := range localAddress {
			if target == addr {
				return errors.New("This is local ip: " + target)
			}
		}
		return nil
	}
	return errors.New("Wrong ip: " + target)
}
