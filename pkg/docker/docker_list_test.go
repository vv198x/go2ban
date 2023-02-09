package docker

import (
	"fmt"
	"github.com/docker/docker/client"
	"testing"
)

type mockOsUtil struct{}

func (m mockOsUtil) CheckFile(file string) bool {
	if file == "/var/lib/docker/containers/container1/container1-json.log" {
		return true
	}
	if file == "/var/lib/docker/containers/container2/container2-json.log" {
		return false
	}
	return false
}

func TestGetListsSyslogFiles(t *testing.T) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		t.Fatalf("error creating client: %v", err)
	}
	defer cli.Close()

	_, err = GetListsSyslogFiles()
	if err != nil {
		fmt.Println("Docker not install")
	}
}
