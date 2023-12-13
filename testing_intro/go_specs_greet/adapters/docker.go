package adapters

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/alecthomas/assert"
	"github.com/docker/go-connections/nat"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func StartDockerServer(t testing.TB, port, dockerFilePath string, binToBuild string) {
	ctx := context.Background()

	req := testcontainers.ContainerRequest{
		FromDockerfile: testcontainers.FromDockerfile{
			Context:    "../../.",
			Dockerfile: dockerFilePath,
			BuildArgs: map[string]*string{
				"bin_to_build": &binToBuild,
			},
			PrintBuildLog: true,
		},
		ExposedPorts: []string{fmt.Sprintf("%s:%s", port, port)},
		WaitingFor:   wait.ForListeningPort(nat.Port(port)).WithStartupTimeout(5 * time.Second),
	}

	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})

	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		assert.NoError(t, container.Terminate(ctx))
	})
}
