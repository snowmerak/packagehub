package windows

import (
	"fmt"
	"os"
	"os/exec"
)

func SetX(key, value string) error {
	cmd := exec.Command("setx", key, value)
	cmd.Stdin = os.Stdin
	result, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("SetX: %w", err)
	}
	fmt.Println(string(result))
	return nil
}

func GetEnv(key string) string {
	return os.Getenv(key)
}
