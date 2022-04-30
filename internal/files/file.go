package files

import (
	"fmt"
	"io/ioutil"
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
