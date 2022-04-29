package files

import (
	"fmt"
	"io/fs"
	"path"
	"path/filepath"
	"strings"
	"webdiff/internal/config"
)

type FileEntries []FileEntry
type FileEntry struct {
	File    string
	Session string
}

func Files(folder string) (FileEntries, error) {
	var (
		rootDir = config.DataDir()
		files   FileEntries
	)

	if len(folder) > 0 {
		rootDir = path.Join(rootDir, folder)
	}

	err := filepath.Walk(rootDir, func(p string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		if strings.Contains(info.Name(), ".response.json") {
			return nil
		}

		fileDir := filepath.Base(filepath.Dir(p))
		files = append(files, FileEntry{
			File:    info.Name(),
			Session: fileDir,
		})

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("cannot read files: %w", err)
	}

	return files, nil
}
