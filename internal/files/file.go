package files

import (
	"fmt"
	"io/ioutil"
	"path"
	"webdiff/internal/config"
)

type ArchivedFile struct {
	Name    string
	Session string
	Content []byte
}

func File(session, name string) (*ArchivedFile, error) {
	filePath := path.Join(config.DataDir(), session, name)
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("cannot read archived file '%s' from session '%s': %v", name, session, err)
	}

	return &ArchivedFile{
		Name:    name,
		Session: session,
		Content: content,
	}, nil
}
