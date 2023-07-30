package utils

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"github.com/justinj96/RapidDocker/internal/dbconfig"
)

func CreateDockerClient() *client.Client {
	// Set the desired API version using the client.WithVersion option
	opts := []client.Opt{
		client.FromEnv,
		client.WithVersion("1.41"),
	}

	// Create the Docker client with the specified options
	cli, err := client.NewClientWithOpts(opts...)
	if err != nil {
		panic(err)
	}

	// Test the client by getting the Docker server version
	ctx := context.Background()
	serverVersion, err := cli.ServerVersion(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to Docker server. Version:", serverVersion.Version)

	return cli
}

func PullDockerImage(cli *client.Client, imageName string, imageTag string) {
	ctx, cancel := context.WithTimeout(context.Background(), 240*time.Second)
	defer cancel()

	reader, err := cli.ImagePull(ctx, imageName, types.ImagePullOptions{})

	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, reader)
}

func RunDockerContainer(cli *client.Client, config dbconfig.Configuration) string {
	ctx := context.Background()
	imageName := config.ImageName
	hostIp := config.HostIp
	port := config.Params["port"]
	containerName := config.Params["containerName"]

	config.EnvVarSetup(&config)

	containerConfig := container.Config{
		Image: imageName,
		Env:   config.EnvironmentVariables,
	}
	hostConfig := container.HostConfig{
		PortBindings: nat.PortMap{
			nat.Port(port + "/tcp"): []nat.PortBinding{
				{
					HostIP:   hostIp,
					HostPort: port,
				},
			},
		},
	}

	resp, err := cli.ContainerCreate(ctx, &containerConfig, &hostConfig, nil, nil, containerName)
	if err != nil {
		panic(err)
	}
	containerID := resp.ID

	if err := cli.ContainerStart(ctx, containerID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}
	return containerID
}
