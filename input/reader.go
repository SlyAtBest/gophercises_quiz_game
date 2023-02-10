package input

import (
	"fmt"
	"os"
)

type FileReader struct {
	path string
}

func NewReader(path string) *FileReader {
	return &FileReader{
		path: path,
	}
}

// Used to load the csv file from disk
func (r *FileReader) ReadData() (string, error) {
	data, err := os.ReadFile(r.path)
	if err != nil {
		return "", fmt.Errorf("failed to load file: %v", r.path)
	}

	return string(data), nil
}
