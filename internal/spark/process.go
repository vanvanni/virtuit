package spark

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/vanvanni/virtuit/internal/out"
)

func StartFirecracker(id string) (*os.Process, string, error) {
	socketPath := filepath.Join("/var/virtuit", id+".socket")
	logPath := filepath.Join("/var/log/virtuit", id+".log")

	cmd := exec.Command("firecracker", "--api-sock", socketPath)

	logFile, err := os.Create(logPath)
	if err != nil {
		return nil, "", fmt.Errorf("create log file: %w", err)
	}
	cmd.Stdout = logFile
	cmd.Stderr = logFile

	err = cmd.Start()
	if err != nil {
		return nil, "", fmt.Errorf("start firecracker: %w", err)
	}

	out.Logger.Info(fmt.Sprintf("Spark started with PID %d, log at %s\n", cmd.Process.Pid, logPath))
	return cmd.Process, socketPath, nil
}
