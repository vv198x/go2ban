package validator

import (
	"errors"
	"github.com/vv198x/go2ban/config"
	"github.com/vv198x/go2ban/pkg/osUtil"
	"net"
	"regexp"
)

func CheckIp(target string) (end string, err error) {
	//The first search option in the ip address line
	//target = regexp.MustCompile(`((25[0-5]|(2[0-4]|1\d|[1-9]|)\d)\.?\b){4}`).FindString(target)
	target = FindIpOnByte(target)

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
	return "", errors.New("wrong ip")
}

/*
This is a FindIpOnByte function that tries to find an IP address in a given string and returns it as a string.
The function uses an array of bytes (buf) to store the IP address while parsing it and
several variables (C, I, P) to keep track of the current position in the byte array and in the input string.
The function iterates over the input string, character by character, and uses a series of conditional statements,
to check if the current character is a valid part of the IP address. If a valid character is found,
it is added to the byte array (buf) and the variables are updated accordingly.
If an invalid character is found, the function checks to see if the full IP address has been found (checking whether
whether the number of dots is 3 and the number of allowed characters is greater than 5),
and if so, it uses net.ParseIP to parse the IP address and convert it to the net.IP type.
The .String() method is then used to return the IP address as a string.
The function returns an empty string if no valid IP address is found in the input string.
*/
func FindIpOnByte(sts string) string {
	buf := make([]byte, 16)
	var C, I, P int8
	// Collecting the buffer
	for _, b := range []byte(sts + " ") { // Space for last iteration
		point := (b == '.')
		if ('0' <= b && b <= '9') && (C < 15) || point {
			if point {
				P++
				I = 0
			}
			if I <= 3 { // Block of 3 numbers
				if !(C == 0 && point) { // Not the first point
					buf[C] = b
				} else { // "Wrong entry" rewrite
					C = -1
					P = 0
				}
			} else { // "Wrong entry" rewrite
				C = -1
				P = 0
			}

			C++ // Iterate position in buffer
			I++ // Iterate position in block of 3 numbers
		} else {
			if (P == 3) && (C > 5) {
				//ParseIP results in four bytes and checks - !(n > 0xFF)
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
