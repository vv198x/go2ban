package validator

import (
	"errors"
	"go2ban/pkg/osUtil"
	"regexp"
)

func CheckIp(target string) (end string, err error) {
	localAddress := osUtil.GetLocalIPs()
	target = regexp.MustCompile(`^((25[0-5]|(2[0-4]|1\d|[1-9]|)\d)\.?\b){4}`).FindString(target)
	if target != "" {
		for _, addr := range localAddress {
			if target == addr {
				return "", errors.New("This is local ip: " + target)
			}
		}
		return target, nil
	}
	return "", errors.New("Wrong ip: " + target)
}
