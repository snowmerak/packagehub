package windows

import (
	"fmt"
	"strings"

	"github.com/luisiturrios/gowin"
)

func GetUserPathEnv() (string, error) {
	val, err := gowin.GetReg("HKCU", "Environment", "Path")
	if err != nil {
		return "", fmt.Errorf("GetUserPathEnv: gowin.GetReg: %v", err)
	}
	return val, nil
}

func AppendUserPathEnv(path string) error {
	prev, err := GetUserPathEnv()
	if err != nil {
		return fmt.Errorf("AppendUserPathEnv: GetUserPathEnv: %v", err)
	}
	if err := gowin.WriteStringReg("HKCU", "Environment", "Path", prev+";"+path); err != nil {
		return fmt.Errorf("AppendUserPathEnv: gowin.WriteStringReg: %v", err)
	}
	return nil
}

func RemoveUserPathEnv(path string) error {
	prev, err := GetUserPathEnv()
	if err != nil {
		return fmt.Errorf("RemoveUserPathEnv: GetUserPathEnv: %v", err)
	}
	prev = strings.ReplaceAll(prev, path, "")
	prev = strings.ReplaceAll(prev, ";;", ";")
	if err := gowin.WriteStringReg("HKCU", "Environment", "Path", prev); err != nil {
		return fmt.Errorf("RemoveUserPathEnv: gowin.WriteStringReg: %v", err)
	}
	return nil
}
