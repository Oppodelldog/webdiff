package files

import (
	"fmt"
)

type FileEntries []FileEntry
type FileEntry struct {
	File    string
	Session string
}

func Files(session string) (FileEntries, error) {
	var (
		files FileEntries
	)

	err := IterateDownloadedFiles(session, func(file FileEntry) error {
		files = append(files, file)
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("cannot iterate downloaded files in session '%s': %w", session, err)
	}

	return files, nil
}
