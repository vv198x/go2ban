package firewall

import (
	"bytes"
	"context"
	"fmt"
	"github.com/vv198x/go2ban/config"
	"log"
	"math/rand"
	"os"
	"strings"
	"testing"
	"time"
)

var fw = &iptables{}

func Test_iptables_Worker(t *testing.T) {
	tests := []struct {
		name    string
		blocked int
	}{
		{"Test create chan", 0},
		{"Worker del ", 10},
	}
	fw.Worker()
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config.Get().BlockedIps = tt.blocked
			for j := 0; j <= tt.blocked*2; j++ {
				fw.Block(context.Background(), getRandomIP())
			}
			fw.Worker()
			// Wait worker clear
			time.Sleep(time.Millisecond * 200)
			if i == 0 {
				if !findIptales("j go2ban") {
					t.Errorf("Chan go2ban not found")
				}
			}
			if i == 1 {
				// (count-cfgMaxLocked)+cfgMaxLocked/10
				// The worker must remove 11 out of 20

				if fw.countBlocked() < 9 {
					t.Errorf("Worker doesn't count correctly")
				}
			}
		})
	}
}

func Test_iptables_Block(t *testing.T) {
	type args struct {
		ctx context.Context
		ip  string
	}
	tests := []struct {
		name string
		args args
	}{
		{"Valid ip", args{context.Background(), "192.168.66.1"}},
		{"Wrong ip", args{context.Background(), "vds"}},
	}
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer log.SetOutput(os.Stderr)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fw.Block(tt.args.ctx, tt.args.ip)
		})
	}

	if !findIptales(tests[0].args.ip) {
		t.Errorf("Iptables not block ip")
	}
	if !strings.Contains(buf.String(), "Not blocked  vds") {
		t.Errorf("Iptables not error")
	}
}

func findIptales(s string) bool {
	byt, err := runOutputCMD("iptables-save")
	if err == nil {
		return bytes.Contains(byt, []byte(s))
	}
	return false
}

func getRandomIP() string {
	rand.Seed(time.Now().UTC().UnixNano())
	ip := fmt.Sprintf("%d.%d.%d.%d",
		1+rand.Intn(254-1), 0+rand.Intn(255-0), 0+rand.Intn(255-0), 1+rand.Intn(253-1))
	return ip
}

func Test_iptables_countBlocked(t *testing.T) {
	tests := []struct {
		name    string
		wantIps int
	}{
		{"add five", 5},
		{"add ten", 10},
	}
	if _, err := fw.UnlockAll(context.Background()); err != nil {
		t.Errorf("Don't clear chan")
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for j := 0; j < tt.wantIps; j++ {
				fw.Block(context.Background(), getRandomIP())
			}

			if gotIps := fw.countBlocked(); gotIps != tt.wantIps {
				t.Errorf("countBlocked() = %v, want %v", gotIps, tt.wantIps)
			}
			if _, err := fw.UnlockAll(context.Background()); err != nil {
				t.Errorf("Don't clear chan")
			}
		})
	}
}

func Test_iptables_GiveBlocked(t *testing.T) {
	t.Run("Check end", func(t *testing.T) {
		fw.Block(context.Background(), "123.123.123.123")
		ips := fw.GetBlocked()
		if _, find := ips["123.123.123.123"]; !find {
			t.Errorf("GetBlocked wrong ip")
		}
	})

}

func Test_iptables_UnlockAll(t *testing.T) {
	fw.Worker()
	time.Sleep(time.Millisecond * 100)

	tests := []struct {
		name    string
		wantErr bool
	}{
		{"All OK", false},
		//{"Delete chan go2ban", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			/*
				if i == 1 {
					if err := runCMD("iptables --table raw --delete PREROUTING --jump go2ban"); err != nil {
						t.Errorf("runCMD error %v", err)
					}
					time.Sleep(time.Millisecond * 100)
					if err := runCMD("iptables --table raw --delete-chain go2ban"); err != nil {
						t.Errorf("runCMD error %v", err)
					}
				}
			*/
			_, err := fw.UnlockAll(context.Background())
			if (err != nil) != tt.wantErr {
				t.Errorf("UnlockAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}
