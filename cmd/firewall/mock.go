package firewall

import (
	"context"
	"errors"
	"log"
)

type Mock struct{}

func (fw *Mock) Block(ctx context.Context, ip string) {
	log.Println("Mock firewall blocked ip:", ip)
}
func (fw *Mock) Worker() {
	log.Println("Mock firewall worker")
}
func (fw *Mock) UnlockAll(ctx context.Context) (ips int, err error) {
	return 0, errors.New("Mock err")
}
func (fw *Mock) countBlocked() (ips int) {
	return 0
}
