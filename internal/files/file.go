package files

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/andybalholm/cascadia"
	"golang.org/x/net/html"
	"io/ioutil"
)

const filterTypeCSS = "css"

type ArchivedFile struct {
	Id      string
	Session string
	Content []byte
}

var ErrFilterNoMatch = errors.New("filter did not match, filtered result was empty")
var ErrFilterInvalid = errors.New("filter is not valid")
var ErrParsingFailed = errors.New("parsing failed")

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

		err = FileFilter(file, filter.Def, filter.Type)
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

func FileFilter(file *ArchivedFile, filter, filterType string) error {
	var err error

	switch filterType {
	case filterTypeCSS:
		file.Content, err = filterCSS(file.Content, filter)
	}

	return err
}

func filterCSS(content []byte, selector string) ([]byte, error) {
	s, err := cascadia.Compile(selector)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrFilterInvalid, err)
	}

	doc, err := html.Parse(bytes.NewBuffer(content))
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrParsingFailed, err)
	}

	var newContent = bytes.NewBuffer(nil)

	nodes := s.MatchAll(doc)
	if len(nodes) == 0 {
		return nil, ErrFilterNoMatch
	}

	for _, node := range nodes {
		if err := html.Render(newContent, node); err != nil {
			return nil, err
		}
	}

	return newContent.Bytes(), nil
}
