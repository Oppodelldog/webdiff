package files

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"webdiff/internal/config"
)

func Sessions() ([]string, error) {
	var (
		dataDir  = config.DataDir()
		sessions []string
	)

	entries, err := os.ReadDir(dataDir)
	if err != nil {
		return nil, fmt.Errorf("cannot read sessions: %w", err)
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		sessions = append(sessions, entry.Name())
	}

	return sessions, nil
}

func SessionUrls(name string) ([]string, error) {
	sessionPath := SessionPath(name)
	entries, err := os.ReadDir(sessionPath)
	if err != nil {
		return nil, fmt.Errorf("cannot read sessions")
	}

	var urls []string
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		if filepath.Ext(entry.Name()) != ".json" {
			continue
		}

		statusFilePath := path.Join(sessionPath, entry.Name())
		f, err := os.Open(statusFilePath)
		if err != nil {
			fmt.Printf("cannot read session file '%s'", statusFilePath)

			continue
		}
		var statusFile PageDownloadStatus
		err = json.NewDecoder(f).Decode(&statusFile)
		if err != nil {
			fmt.Printf("cannot decode session file '%s'", entry.Name())

			continue
		}

		urls = append(urls, statusFile.Request.URL.String())

	}

	return urls, nil
}
