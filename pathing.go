package prefer

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
)

var standardPaths []string

func getWindowsPaths() []string {
	return []string{
		os.Getenv("USERPROFILE"),
		os.Getenv("LOCALPROFILE"),
		os.Getenv("APPDATA"),

		os.Getenv("CommonProgramFiles"),

		os.Getenv("ProgramData"),
		os.Getenv("ProgramFiles"),
		os.Getenv("ProgramFiles(x86)"),

		os.Getenv("SystemRoot"),
		filepath.Join(os.Getenv("SystemRoot"), "system32"),
	}
}

func getUnixPaths() []string {
	return []string{
		filepath.Join(os.Getenv("HOME"), ".config"),
		os.Getenv("HOME"),
		"/usr/local",
		"/usr",
		"/",
	}
}

func StandardPaths() []string {
	paths := make([]string, len(standardPaths))
	copy(paths, standardPaths)
	return paths
}

func distinct(items []string) (result []string) {
Search:
	for _, item := range items {
		for _, existing := range result {
			if existing == item {
				continue Search
			}
		}

		result = append(result, item)
	}

	return result
}

func init() {
	wd, err := os.Getwd()

	if err != nil {
		log.Fatalln("Can not get current working directory.")
	}

	// Remove /bin if it's at the end of the cwd
	// TODO: Er, os.PathSeparator is a rune... So, here's a hack.
	if len(wd) > 4 && wd[len(wd)-4:] == filepath.Join("", "bin") {
		wd = wd[:len(wd)-4]
	}

	paths := []string{".", wd}
	xdgPaths := filepath.SplitList(os.Getenv("XDG_CONFIG_DIRS"))

	if len(xdgPaths) > 0 {
		paths = append(paths, xdgPaths...)
	}

	switch runtime.GOOS {
	case "win32":
		paths = append(paths, getWindowsPaths()...)
	default:
		paths = append(paths, getUnixPaths()...)
	}

	for _, path := range paths {
		if path == "/" {
			path = ""
		}

		standardPaths = append(standardPaths, filepath.Join(path, "/etc"))

		if path == "" {
			continue
		}

		if len(path) > 0 {
			standardPaths = append(standardPaths, path)
		}
	}

	standardPaths = distinct(standardPaths)
}
