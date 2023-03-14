package abuseipdb

import (
	"fmt"
	"github.com/vv198x/go2ban/cmd/firewall"
	"github.com/vv198x/go2ban/config"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestScheduler(t *testing.T) {
	firewall.ExportFirewall = &firewall.Mock{}
	config.Get().Flags.RunAsDaemon = true
	t.Run("empty api key", func(t *testing.T) {
		go Scheduler("")
	})
	t.Run("not empty api key", func(t *testing.T) {
		go Scheduler("test")
	})
}

func TestBlockBlackListIPs(t *testing.T) {
	firewall.ExportFirewall = &firewall.Mock{}
	// Set up a mock HTTP server to return a test blacklist
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "1.2.3.4")
		fmt.Fprintln(w, "5.6.7.8")
		t.Run("api raw string", func(t *testing.T) {
			wantSt := "/?confidenceMinimum=90&limit=0"
			if r.URL.String() != wantSt {
				t.Errorf("blockBlackListIPs send:%s, want:%s", r.URL.String(), wantSt)
			}
		})
		t.Run("Check heandlers", func(t *testing.T) {
			wantKey := "test-api-key"
			wantAccept := "text/plain"
			if r.Header.Get("Key") != wantKey {
				t.Errorf("blockBlackListIPs send:%s, want:%s", r.Header.Get("Key"), wantKey)
			}
			if r.Header.Get("Accept") != wantAccept {
				t.Errorf("blockBlackListIPs send:%s, want:%s", r.Header.Get("Accept"), wantAccept)
			}
		})
	}))
	defer mockServer.Close()

	// Call the blockBlackListIPs function with the mock server and firewall
	blockBlackListIPs("test-api-key", mockServer.URL)

}
