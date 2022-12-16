package validator

import (
	"fmt"
	"testing"
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
	ip, err := CheckIp("212.354.254.0")
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

func TestGoodIPendText(t *testing.T) {
	ip, err := CheckIp("212.254.254.254TEXT")
	if ip != "212.254.254.254" {
		t.Errorf("\nReturn ip:%s, err:%s", ip, err)
	} else {
		fmt.Println("Validator say error: ", err)
		fmt.Println("Validator say ip: ", ip)
	}
}
