package docker

import (
	"fmt"
	"github.com/docker/docker/client"
	"testing"
)

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
