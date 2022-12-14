package validator

import (
	"errors"
	"go2ban/pkg/config"
	"go2ban/pkg/osUtil"
	"regexp"
)

func CheckIp(target string) (end string, err error) {
	target = regexp.MustCompile(`((25[0-5]|(2[0-4]|1\d|[1-9]|)\d)\.?\b){4}`).FindString(target)
	if target != "" {
		whiteAddress := osUtil.GetLocalIPs()
		whiteAddress = append(whiteAddress, config.Get().WhiteList...)
		err = errors.New("This is white ip: " + target)
		for _, addr := range whiteAddress {
			if addr[len(addr)-1] == '*' {
				rexp := regexp.MustCompile(`^((25[0-5]|(2[0-4]|1\d|[1-9]|)\d)\.?){3}`)
				if rexp.FindString(addr) == rexp.FindString(target) {
					return "", err
				}
			}
			if target == addr {
				return "", err
			}
		}
		return target, nil
	}
	return "", errors.New("Wrong ip")
}
