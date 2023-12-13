package main_test

import (
	"fmt"
	"go_specs_greet/adapters"
	"go_specs_greet/adapters/grpcserver"
	"go_specs_greet/specifications"
	"testing"
)

func TestGreeterServer(t *testing.T) {
	var (
		port           = "50051"
		dockerFilePath = "./Dockerfile"
		driver         = grpcserver.Driver{Addr: fmt.Sprintf("localhost:%s", port)}
	)

	adapters.StartDockerServer(t, port, dockerFilePath, "grpcserver")
	specifications.GreetSpecification(t, &driver)
	specifications.CurseSepcification(t, &driver)
}
