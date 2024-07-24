package util

import "os"

func GetFileContent(path string) (string, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
