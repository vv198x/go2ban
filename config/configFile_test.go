package config

import (
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	testCfg := "/tmp/test_config.config"
	exportCfg.Flags.ConfigFile = testCfg
	if exportCfg.Flags.ConfigFile != testCfg {
		t.Fatalf("Incorrect flag value for ConfigFile")
	}

	// Create a test configuration file with a variety of parameters and JSON data.
	testConfigData := `
firewall=iptables #comment

# This is a comment
log_dir=/tmp/log/go2ban #comment
white_list=192.168.0.1 192.168.0.* #comment

grpc_port=111/tcp #comment
blocked_ips=9000 #comment
fake_socks_ports=22 21 3389 #comment
fake_socks_fails=2 #comment
rest_port=222 #comment
local_service_check_minutes=2 #comment
local_service_fails=2 #comment
{
  "Service":[
    {"On":true,"Name":"sshd_cent","Regxp": "Failed password","LogFile":"/var/log/secure"},
    {"On":false,"Name":"postree11_local","Regxp": "password authentication failed","LogFile":"/var/log/go2ban/test.log"},
    {"On":true,"Name":"postree11_docker","Regxp": "password authentication failed","LogFile":"docker"},
    {"On":true,"Name":"shandow_socks","Regxp": "authentication error","LogFile":"docker"}
  ]
}`
	file, err := os.Create(exportCfg.Flags.ConfigFile)
	if err != nil {
		t.Fatalf("Error creating test config file: %v", err)
	}
	defer file.Close()

	_, err = file.Write([]byte(testConfigData))
	if err != nil {
		t.Fatalf("Error writing test config file: %v", err)
	}

	err = file.Sync()
	if err != nil {
		t.Fatalf("Error syncing test config file: %v", err)
	}

	// Read file
	Load()

	if exportCfg.GrpcPort != "111/tcp" {
		t.Fatalf("Incorrect GrpcPort value")
	}

	if exportCfg.RestPort != "222" {
		t.Fatalf("Incorrect RestPort value")
	}

	if exportCfg.Firewall != "iptables" {
		t.Fatalf("Incorrect Firewall value")
	}

	if exportCfg.LogDir != "/tmp/log/go2ban" {
		t.Fatalf("Incorrect LogDir value")
	}

	if exportCfg.WhiteList[1] != "192.168.0.*" {
		t.Fatalf("Incorrect WhiteList value")
	}

	if exportCfg.BlockedIps != 9000 {
		t.Fatalf("Incorrect BlockedIps value")
	}

	if exportCfg.FakeSocksPorts[2] != 3389 {
		t.Fatalf("Incorrect FakeSocksPorts value")
	}

	if exportCfg.FakeSocksFails != 2 {
		t.Fatalf("Incorrect FakeSocksFails value")
	}

	if exportCfg.ServiceCheckMinutes != 2 {
		t.Fatalf("Incorrect ServiceCheckMinutes value")
	}

	if exportCfg.ServiceFails != 2 {
		t.Fatalf("Incorrect ServiceFails value")
	}

	if exportCfg.Services[1].On {
		t.Fatalf("Incorrect Service.On value")
	}

	for _, s := range exportCfg.Services {
		if (s.Name == "") && (s.LogFile == "") && (s.Regxp == "") {
			t.Fatalf("Incorrect Service value")
		}
	}

	err = os.Remove(exportCfg.Flags.ConfigFile)
	if err != nil {
		t.Fatalf("Failed to remove file: %v", err)
	}
}

func Test_whatFirewall(t *testing.T) {
	tests := []struct {
		name             string
		wantFirewallType string
		wantErr          bool
	}{
		{
			name:             "iptables found",
			wantFirewallType: "iptables",
			wantErr:          false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil && !tt.wantErr {
					t.Errorf("whatFirewall() panicked = %v, wantErr %v", r, tt.wantErr)
				}
			}()
			if gotFirewallType := whatFirewall(); gotFirewallType != tt.wantFirewallType {
				t.Errorf("whatFirewall() = %v, want %v", gotFirewallType, tt.wantFirewallType)
			}
		})
	}
}
