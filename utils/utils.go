package utils

import (
	"io"
	"os"
)

func ReadContents(filename string) (string, error) {
	file, err := os.Open(filename)
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	contents, err := io.ReadAll(file)
	return string(contents), err
}
