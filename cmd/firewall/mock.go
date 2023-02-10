package firewall

import "context"

type Mock struct{}

func (fw *Mock) Block(ctx context.Context, ip string) {}
func (fw *Mock) Worker()                              {}
func (fw *Mock) UnlockAll(ctx context.Context) (ips int, err error) {
	return 0, err
}
func (fw *Mock) countBlocked() (ips int) {
	return 0
}
