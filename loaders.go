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
	Discover(identifier string) (string, error)
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

func discoverExtensionForPath(path string) (string, error) {
	for extension := range serializers {
		fullPath := path + "." + extension

		if stat, err := os.Stat(fullPath); !os.IsNotExist(err) {
			if stat.IsDir() {
				continue
			}

			return fullPath, nil
		}
	}

	return "", fmt.Errorf("could not find any matching files in %v", path)
}

func discoverFileInPath(dir, identifier string) (string, error) {
	fullPath := path.Join(dir, identifier)

	if extension := filepath.Ext(fullPath); extension == "" {
		return discoverExtensionForPath(fullPath)
	} else if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		return "", fmt.Errorf("unable to find matching configuration in %v", dir)
	}

	return fullPath, nil
}

// Discover attempts to match the given identifier to the closest file match
func (loader FileLoader) Discover(identifier string) (string, error) {
	var result string
	var err error

	for _, currentPath := range StandardPaths() {
		if currentPath, err = discoverFileInPath(currentPath, identifier); err != nil {
			continue
		}

		result = currentPath
		break
	}

	return result, err
}

// Load loads a specific file from the filesystem
func (loader FileLoader) Load(identifier string) ([]byte, string, error) {
	var result []byte

	file, err := os.Open(identifier)
	panicIfError(err)
	defer file.Close()

	for scanner := bufio.NewScanner(file); scanner.Scan(); {
		result = append(result, scanner.Bytes()...)
	}

	return result, identifier, err
}
