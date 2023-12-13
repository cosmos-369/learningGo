package main

import (
	"fmt"
	"math/rand"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
	"time"
)

const baseBinName = "temp-testbinary"

func LaunchTestProgram(port string) (cleanup func(), sendInterupt func() error, err error) {
	binName, err := buildBinary()
	if err != nil {
		return nil, nil, err
	}

	sendInterupt, kill, err := runServer(binName, port)

	//cleanup fuction delets the binary file created
	cleanup = func() {
		if kill != nil {
			kill()
		}
		os.Remove(binName)
	}

	//if you get an error
	//clean up the files
	if err != nil {
		cleanup()
		return nil, nil, err
	}

	return cleanup, sendInterupt, nil
}

func buildBinary() (binName string, err error) {
	//generate a random name for the bin file
	binName = randomString(10) + "-" + baseBinName

	build := exec.Command("go", "build", "-o", binName)

	if err := build.Run(); err != nil {
		return "", fmt.Errorf("cannot build tool %s: %s", binName, err)
	}

	return binName, nil
}

func runServer(binName string, port string) (sendInterrupt func() error, kill func(), err error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, nil, err
	}

	cmdPath := filepath.Join(dir, binName)
	cmd := exec.Command(cmdPath)
	//run the binary file
	if err := cmd.Start(); err != nil {
		return nil, nil, fmt.Errorf("cannot run temp converter: %s", err)
	}

	//grab the function to kill the process which is running the server(biary file)
	kill = func() {
		_ = cmd.Process.Kill()
	}

	//send SIGTERM to the process (server)
	sendInterrupt = func() error {
		return cmd.Process.Signal(syscall.SIGTERM)
	}

	err = waitForServerListening(port)

	return
}

func waitForServerListening(port string) error {

	//tries to connect to the server
	for i := 0; i < 30; i++ {
		conn, _ := net.Dial("tcp", net.JoinHostPort("localhost", port))
		if conn != nil {
			conn.Close()
			return nil
		}
		time.Sleep(100 * time.Millisecond)
	}
	return fmt.Errorf("nothing seems to be listening on localhost:%s", port)
}

func randomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}

	return string(s)
}
