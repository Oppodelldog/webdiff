package files

import (
	"fmt"
	"os"
	"path"
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

	sessionDirs, err := os.ReadDir(rootDir)
	if err != nil {
		return nil, fmt.Errorf("cannot read session dirs: %w", err)
	}
	for _, sessionDir := range sessionDirs {
		if len(folder) > 0 && sessionDir.Name() == folder {
			continue
		}

		downloadFiles, err := os.ReadDir(path.Join(rootDir, sessionDir.Name(), archiveSubFolder))
		if err != nil {
			continue
		}

		for _, file := range downloadFiles {
			files = append(files, FileEntry{
				File:    file.Name(),
				Session: sessionDir.Name(),
			})
		}
	}

	return files, nil
}
