package prefer

import (
	"bufio"
	"os"
)

type Loader interface {
	Load(identifier string) ([]byte, error)
}

func NewLoader(identifier string) (loader Loader, err error) {
	switch identifier {
	default:
		return loader, nil
	}
}

type FileLoader struct{}

func (loader FileLoader) Load(identifier string) (result []byte, err error) {
	file, err := os.Open(identifier)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, scanner.Bytes()...)
	}

	return result, err
}
