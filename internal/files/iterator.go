package files

import (
	"encoding/json"
	"fmt"
	"github.com/Oppodelldog/webdiff/internal/config"
	"io/fs"
	"os"
	"path"
	"path/filepath"
)

func IterateDownloadedFiles(session string, processDownloadedFile func(FileEntry) error) error {
	return IterateSessionDirs(session, func(sessionDirName string) error {
		return IterateDirFiles(DownloadedDirPath(sessionDirName), func(fileName string) error {
			err := processDownloadedFile(FileEntry{
				File:    fileName,
				Session: sessionDirName,
			})
			if err != nil {
				return fmt.Errorf("cannot read session downloaded file '%s': %w", fileName, err)
			}

			return nil
		})
	})
}

func IterateStatusFiles(session string, processStatusFile func(PageDownloadStatus) error) error {
	return IterateSessionDirs(session, func(sessionDirName string) error {
		return IterateDirFiles(SessionPath(session), func(fileName string) error {
			if filepath.Ext(fileName) != ".json" {
				return nil
			}

			statusFilePath := path.Join(SessionPath(sessionDirName), fileName)
			f, err := os.Open(statusFilePath)
			if err != nil {
				return fmt.Errorf("cannot read session file '%s'", statusFilePath)
			}

			var statusFile PageDownloadStatus
			err = json.NewDecoder(f).Decode(&statusFile)
			if err != nil {
				return fmt.Errorf("cannot decode session file '%s'", statusFilePath)
			}

			err = processStatusFile(statusFile)
			if err != nil {
				return fmt.Errorf("cannot read session downloaded file '%s': %w", fileName, err)
			}

			return nil
		})
	})
}

func IterateSessionDirs(session string, processSessionDir func(string) error) error {
	return Iterate(config.DataDir(), func(entry fs.DirEntry) error {
		if !entry.IsDir() {
			return nil
		}

		if len(session) > 0 && entry.Name() != session {
			return nil
		}

		err := processSessionDir(entry.Name())
		if err != nil {
			fmt.Printf("cannot process session dir '%s': %v", entry.Name(), err)

			return nil
		}

		return nil
	})
}

func IterateDirFiles(dir string, processFile func(string) error) error {
	return Iterate(dir, func(entry fs.DirEntry) error {
		if entry.IsDir() {
			return nil
		}

		err := processFile(entry.Name())
		if err != nil {
			return fmt.Errorf("cannot process file '%s': %w", entry.Name(), err)
		}

		return nil
	})
}

func Iterate(dir string, processDirEntry func(entry fs.DirEntry) error) error {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return fmt.Errorf("cannot read dir '%s': %w", dir, err)
	}

	for _, entry := range entries {
		err := processDirEntry(entry)
		if err != nil {
			return fmt.Errorf("cannot process dir entries '%s': %w", entry.Name(), err)
		}
	}

	return nil
}
