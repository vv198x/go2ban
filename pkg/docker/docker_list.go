package docker

import (
	"context"
	"github.com/vv198x/go2ban/pkg/osUtil"
	"path/filepath"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func GetListsSyslogFiles() ([]string, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}
	defer cli.Close()

	//Получаю все контейнеры
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		return nil, err
	}

	//Проверяю суслог файл у контейнера
	syslogPaths := make([]string, 0, len(containers))
	for _, container := range containers {
		syslogPath := filepath.Join("/var/lib/docker/containers", container.ID, container.ID+"-json.log")
		if osUtil.CheckFile(syslogPath) {

			syslogPaths = append(syslogPaths, syslogPath)
		}
	}

	return syslogPaths, nil
}
