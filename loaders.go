package prefer

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"path/filepath"
)

// A Loader represents a specific way of loading bytes from a data source
type Loader interface {
	Discover(identifier string) string
	Load(identifier string) ([]byte, string, error)
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

// Discover attempts to match the given identifier to the closest file match
func (loader FileLoader) Discover(identifier string) string {
	extension := path.Ext(identifier)

	if extension == "" {
		for key := range serializers {
			extension = "." + key

			if _, err := os.Stat(identifier + extension); os.IsNotExist(err) != true {
				identifier += extension
				break
			}

			extension = ""
		}
	}

	if absolute, err := filepath.Abs(identifier); err == nil {
		identifier = absolute
	}

	return identifier
}

// Load loads a specific file from the filesystem
func (loader FileLoader) Load(identifier string) ([]byte, string, error) {
	var result []byte

	if identifier == "" {
		return nil, "", fmt.Errorf("unable to find matching file for the given identifier (%v)", identifier)
	}

	file, err := os.Open(identifier)
	panicIfError(err)
	defer file.Close()

	for scanner := bufio.NewScanner(file); scanner.Scan(); {
		result = append(result, scanner.Bytes()...)
	}

	return result, identifier, err
}
