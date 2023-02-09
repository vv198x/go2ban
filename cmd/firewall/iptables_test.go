package firewall

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestIptablesBlock(t *testing.T) {
	ip := getRandomIP()
	BlockIP(context.Background(), ip)
	//byt, err := runOutputCMD("iptables-save")
	//fmt.Println(string(byt), err)
}
func BenchmarkIptablesBlock(b *testing.B) {
	ctx := context.Background()

	b.Run("100 iterations", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			for i := 0; i < 100; i++ {
				iptablesBlock(ctx, getRandomIP())
			}
		}
	})

	b.Run("500 iterations", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			for i := 0; i < 900; i++ {
				iptablesBlock(ctx, getRandomIP())
			}
		}
	})

	b.Run("2000 iterations", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			for i := 0; i < 2000; i++ {
				iptablesBlock(ctx, getRandomIP())
			}
		}
	})
}

func BenchmarkIptablesUnlockAll(b *testing.B) {
	b.Run("1 iteration", func(b *testing.B) {
		iptablesUnlockAll(context.Background())
	})
}

func getRandomIP() string {
	rand.Seed(time.Now().UTC().UnixNano())
	ip := fmt.Sprintf("%d.%d.%d.%d",
		1+rand.Intn(255-1), 0+rand.Intn(255-0), 0+rand.Intn(255-0), 1+rand.Intn(254-1))
	return ip
}
