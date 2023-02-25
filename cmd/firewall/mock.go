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
func (fw *Mock) GetBlocked() map[string]struct{} {
	m := make(map[string]struct{})
	m["123.123.123.123"] = struct{}{}
	return m
}
