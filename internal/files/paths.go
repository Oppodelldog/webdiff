package files

import (
	"fmt"
	"github.com/Oppodelldog/webdiff/internal/config"
	"path"
)

const archiveSubFolder = "downloaded"
const filtersFile = "filters.json"

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
	return path.Join(DataDir(), session)
}

func FiltersFile() string {
	return path.Join(DataDir(), filtersFile)
}

func DataDir() string {
	return config.DataDir()
}
