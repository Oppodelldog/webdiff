package files

import (
	"fmt"
	"os"
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
