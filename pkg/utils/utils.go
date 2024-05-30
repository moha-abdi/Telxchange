package utils

import (
	"os"
)

func JSONRead(file string) (result map[string]interface{}, err error) {
	fileData, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
}
