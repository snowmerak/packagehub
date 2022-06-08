package windows

import (
	"fmt"
	"os"
	"os/exec"
)

func SetX(key, value string) error {
	cmd := exec.Command("setx", key, value)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("SetX: cmd.Start: %w", err)
	}
	if err := cmd.Wait(); err != nil {
		return fmt.Errorf("SetX: cmd.Wait: %w", err)
	}
	return nil
}

func GetEnv(key string) string {
	return os.Getenv(key)
}

func AddPath(path string) error {
	return SetX("Path", GetEnv("Path")+";"+path)
}
