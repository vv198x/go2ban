package validator

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestGoodIP(t *testing.T) {
	ip, err := CheckIp("127.0.0.2")
	if ip != "127.0.0.2" {
		t.Errorf("\nReturn not 127.0.0.2 :%s, err:%s", ip, err)
		fmt.Println()
	} else {
		fmt.Println("Validator say error: ", err)
		fmt.Println("Validator say ip: ", ip)
	}
}
func TestHardGoodIP(t *testing.T) {
	ip, err := CheckIp("127.0.0.2:")
	if ip != "127.0.0.2" {
		t.Errorf("\nReturn not 127.0.0.2 :%s, err:%s", ip, err)
	} else {
		fmt.Println("Validator say error: ", err)
		fmt.Println("Validator say ip: ", ip)
	}
}

func TestLocalIP(t *testing.T) {
	ip, err := CheckIp("127.0.0.1")
	if ip != "" {
		t.Errorf("\nReturn ip:%s, err:%s", ip, err)
	} else {
		fmt.Println("Validator say error: ", err)
		fmt.Println("Validator say ip: ", ip)
	}
}

func TestNullIP(t *testing.T) {
	ip, err := CheckIp("0.0.0.127")
	if ip != "" {
		t.Errorf("\nReturn ip:%s, err:%s", ip, err)
	} else {
		fmt.Println("Validator say error: ", err)
		fmt.Println("Validator say ip: ", ip)
	}
}

func TestBadlIP212(t *testing.T) {
	ip, err := CheckIp("212.354.254.254")
	if ip != "" {
		t.Errorf("\nReturn ip:%s, err:%s", ip, err)
	} else {
		fmt.Println("Validator say error: ", err)
		fmt.Println("Validator say ip: ", ip)
	}
}

func TestBadlIPZero(t *testing.T) {
	ip, err := CheckIp("212.254.254.0")
	if ip != "" {
		t.Errorf("\nReturn ip:%s, err:%s", ip, err)
	} else {
		fmt.Println("Validator say error: ", err)
		fmt.Println("Validator say ip: ", ip)
	}
}

func TestGoodIPstartText(t *testing.T) {
	ip, err := CheckIp("TEXT212.254.254.254")
	if ip != "212.254.254.254" {
		t.Errorf("\nReturn ip:%s, err:%s", ip, err)
	} else {
		fmt.Println("Validator say error: ", err)
		fmt.Println("Validator say ip: ", ip)
	}
}

func TestBadlIP312(t *testing.T) {
	ip, err := CheckIp("312.354.254.254")
	if ip != "" {
		t.Errorf("\nReturn ip:%s, err:%s", ip, err)
	} else {
		fmt.Println("Validator say error: ", err)
		fmt.Println("Validator say ip: ", ip)
	}
}

func TestFindSSHD(t *testing.T) {
	ip, err := CheckIp(`Dec 14 23:54:25 hostname sshd[14105]: Failed password for root from 123.123.123.123 port 47138 ssh2`)
	if ip != "123.123.123.123" {
		t.Errorf("\nReturn ip:%s, err:%s", ip, err)
	} else {
		fmt.Println("Validator say error: ", err)
		fmt.Println("Validator say ip: ", ip)
	}
}

func TestFindPGsql(t *testing.T) {
	ip, err := CheckIp(`{"log":"2022-12-16 10:17:20.246 UTC [29]   pgsql @pgdb 123.123.123.123 - FATAL:  password authentication failed for user \"pgsql\"\n","stream":"stderr","time":"2022-12-16T10:17:20.247459026Z"}`)
	if ip != "123.123.123.123" {
		t.Errorf("\nReturn ip:%s, err:%s", ip, err)
	} else {
		fmt.Println("Validator say error: ", err)
		fmt.Println("Validator say ip: ", ip)
	}
}

func TestRandom1000(t *testing.T) {
	for i := 0; i < 1000; i++ {
		rand.Seed(time.Now().UTC().UnixNano())
		rndip := fmt.Sprintf("%d.%d.%d.%d",
			1+rand.Intn(255-1), 1+rand.Intn(255-0), 1+rand.Intn(255-0), 1+rand.Intn(255-1))
		//fmt.Println(rndip)
		ip, err := CheckIp(rndip)
		if ip != rndip {
			t.Errorf("\nReturn ip:%s, err:%s", ip, err)
		}
	}
}
