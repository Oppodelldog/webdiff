package files

import (
	"fmt"
	"io/ioutil"
)

type ArchivedFile struct {
	Id      string
	Session string
	Content []byte
}

func FileFiltered(session, id, filterName string) (*ArchivedFile, error) {
	filters, err := Filters()
	if err != nil {
		return nil, fmt.Errorf("filters not accessible")
	}

	file, err := File(session, id)
	if err != nil {
		return nil, fmt.Errorf("file not found: %v", err)
	}

	if len(filterName) > 0 {
		filter := filters.ByName(filterName)
		if filter == nil {
			return nil, fmt.Errorf("filter not found")
		}

		file.Content, err = FilterFile(file.Content, filter.Def, filter.Type)
		if err != nil {
			return nil, fmt.Errorf("error in filter: %w", err)
		}
	}

	return file, nil
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
