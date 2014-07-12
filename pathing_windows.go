package prefer

import (
	"os"
	"path/filepath"
)

func getSystemPaths() []string {
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
