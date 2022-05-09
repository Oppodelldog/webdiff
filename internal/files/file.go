package files

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"log"
)

type ArchivedFile struct {
	Id      string
	Session string
	Content []byte
}

var ErrFilterNoMatch = errors.New("filter did not match, filtered result was empty")

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

		err = FileFilter(file, filter.Filter)
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

func FileFilter(file *ArchivedFile, filter string) error {
	doc, err := goquery.NewDocumentFromReader(bytes.NewBuffer(file.Content))
	if err != nil {
		log.Fatal(err)
	}

	var newContent = bytes.NewBuffer(nil)
	selection := doc.Find(filter)
	if selection.Length() == 0 {
		return ErrFilterNoMatch
	}
	selection.Each(func(i int, selection *goquery.Selection) {
		html, err := goquery.OuterHtml(selection)
		if err != nil {
			fmt.Printf("error selecting html: %v\n", err)
		}
		newContent.WriteString(html)
	})
	if err != nil {
		return fmt.Errorf("cannot filter: %v", err)
	}

	file.Content = newContent.Bytes()

	return nil
}
