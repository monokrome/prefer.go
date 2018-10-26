package prefer

import (
	"bufio"
	"os"
)

// A Loader represents a specific way of loading bytes from a data source
type Loader interface {
	Load(identifier string) ([]byte, error)
}

// NewLoader creates a new Loader based on the given identifier
func NewLoader(identifier string) (Loader, error) {
	switch identifier {
	default:
		return FileLoader{}, nil
	}
}

// FileLoader reads bytes from regular files
type FileLoader struct{}

// Load loads a specific file from the filesystem
func (loader FileLoader) Load(identifier string) (result []byte, err error) {
	file, err := os.Open(identifier)
	panicIfError(err)
	defer file.Close()

	for scanner := bufio.NewScanner(file); scanner.Scan(); {
		result = append(result, scanner.Bytes()...)
	}

	return result, err
}
