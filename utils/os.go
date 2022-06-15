package utils

import (
	"io"
	"os"
	"runtime"
)

func GetOS() string {
	switch runtime.GOOS {
	case "windows":
		return "win64"
	case "linux":
		return "lnx64"
	default:
		return "unknown"
	}
}

func IsDir(path string) bool {
	stat, err := os.Stat(path)
	if err != nil {
		return false
	}
	return stat.IsDir()
}

func IsFile(path string) bool {
	stat, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !stat.IsDir()
}

func CopyFile(src string, dest string) error {
	srcF, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcF.Close()

	destF, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destF.Close()

	_, err = io.Copy(destF, srcF)
	return err
}
