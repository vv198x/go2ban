package firewall

import "context"

type mock struct{}

func (fw *mock) Block(ctx context.Context, ip string) {}
func (fw *mock) Worker()                              {}
func (fw *mock) UnlockAll(ctx context.Context) (ips int, err error) {
	return 0, err
}
func (fw *mock) countBlocked() (ips int) {
	return 0
}
