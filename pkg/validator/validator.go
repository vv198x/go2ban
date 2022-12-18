package validator

import (
	"errors"
	"go2ban/pkg/config"
	"go2ban/pkg/osUtil"
	"net"
	"regexp"
)

func CheckIp(target string) (end string, err error) {

	//target = regexp.MustCompile(`((25[0-5]|(2[0-4]|1\d|[1-9]|)\d)\.?\b){4}`).FindString(target)
	target = FindIpOnByte(target) //20% faster

	if target != "" && target[0] != '0' {
		if target[len(target)-1] == '0' && target[len(target)-2] == '.' {
			return "", err
		}

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

func FindIpOnByte(sts string) string {
	buf := make([]byte, 16)
	var C, I, P int8

	for _, b := range []byte(sts + " ") { // Space for the last iteration
		point := (b == '.')
		if ('0' <= b && b <= '9') && (C < 15) || point {
			if point {
				P++
				I = 0
			}
			if I <= 3 { // in bloks 3 char
				if !(C == 0 && point) { // no first point
					buf[C] = b
				} else {
					C = -1
					P = 0
				}
			} else {
				C = -1
				P = 0
			}

			C++
			I++
		} else {
			if (P == 3) && (C > 5) {

				ip := net.ParseIP(string(buf[:C]))
				if ip != nil {
					return ip.String()
				} else {
					return ""
				}
			}

			C, I, P = 0, 0, 0
		}
	}
	return ""
}
