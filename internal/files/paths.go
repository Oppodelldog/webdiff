package files

import (
	"fmt"
	"path"
	"webdiff/internal/config"
)

func DownloadedFilePath(session, id string) string {
	return path.Join(DownloadedDirPath(session), id)
}
func DownloadedDirPath(session string) string {
	return path.Join(SessionPath(session), archiveSubFolder)
}

func StatusFilePath(session, id string) string {
	return path.Join(SessionPath(session), fmt.Sprintf("%s.json", id))
}

func SessionPath(session string) string {
	return path.Join(config.DataDir(), session)
}
