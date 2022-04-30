package webapp

import (
	"embed"
	"io/fs"
	"os"
)

//go:embed assets
var assetsFS embed.FS

//go:embed templates
var templatesFS embed.FS

// Filesystem returns a fs from the path set in the given env var, otherwise returns the value of parameter fs.
func Filesystem(envVar string, fs fs.FS) fs.FS {
	if absolutePath, ok := os.LookupEnv(envVar); ok {
		return os.DirFS(absolutePath)
	}

	return fs
}
