package files

import (
	"fmt"
	"io/ioutil"
	"path"
	"webdiff/internal/config"
)

const archiveSubFolder = "downloaded"

type ArchivedFile struct {
	Id      string
	Session string
	Content []byte
}

func File(session, id string) (*ArchivedFile, error) {
	filePath := DownloadedFilePath(session, id)
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("cannot read archived file '%s' from session '%s': %v", id, session, err)
	}

	return &ArchivedFile{
		Id:      id,
		Session: session,
		Content: content,
	}, nil
}

func DownloadedFilePath(session, id string) string {
	return path.Join(SessionPath(session), archiveSubFolder, id)
}

func StatusFilePath(session, id string) string {
	return path.Join(SessionPath(session), fmt.Sprintf("%s.json", id))
}

func SessionPath(session string) string {
	return path.Join(config.DataDir(), session)
}
