package prefer

import (
	"bufio"
	"os"
)

type Loader interface {
	Load(identifier string) ([]byte, error)
}

func NewLoader(identifier string) (Loader, error) {
	switch identifier {
	default:
		return FileLoader{}, nil
	}
}

type FileLoader struct{}

func (loader FileLoader) Load(identifier string) (result []byte, err error) {
	file, err := os.Open(identifier)
	check(err)
	defer file.Close()

	for scanner := bufio.NewScanner(file); scanner.Scan(); {
		result = append(result, scanner.Bytes()...)
	}

	return result, err
}
