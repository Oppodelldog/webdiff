package files

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/PaesslerAG/gval"
	"github.com/PaesslerAG/jsonpath"
	"github.com/andybalholm/cascadia"
	"github.com/antchfx/htmlquery"
	"github.com/antchfx/xmlquery"
	"golang.org/x/net/html"
)

const (
	filterTypeCSS       = "css"
	filterTypeHTMLXPATH = "htmlxpath"
	filterTypeXMLXPATH  = "xmlxpath"
	filterTypeJSONPATH  = "jsonxpath"
)

var (
	ErrFilterNoMatch     = errors.New("filter did not match, filtered result was empty")
	ErrFilterInvalid     = errors.New("filter is not valid")
	ErrParsingFailed     = errors.New("parsing failed")
	ErrUnknownFilterType = errors.New("unknown filter type")
)

func FilterFile(content []byte, filter, filterType string) ([]byte, error) {
	switch filterType {
	case filterTypeCSS:
		return filterCSS(content, filter)
	case filterTypeHTMLXPATH:
		return filterHTMLXPath(content, filter)
	case filterTypeXMLXPATH:
		return filterXMLXPath(content, filter)
	case filterTypeJSONPATH:
		return filterJSONXPATH(content, filter)
	}

	return nil, fmt.Errorf("%w '%s'", ErrUnknownFilterType, filterType)
}

func filterJSONXPATH(content []byte, filter string) ([]byte, error) {
	builder := gval.Full(jsonpath.PlaceholderExtension())

	path, err := builder.NewEvaluable(filter)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrFilterInvalid, err)
	}

	var obj interface{}
	err = json.Unmarshal(content, &obj)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrParsingFailed, err)
	}

	nodes, err := path(context.Background(), obj)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrParsingFailed, err)
	}

	jsonOutput, err := json.Marshal(nodes)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrFilterNoMatch, err)
	}

	return jsonOutput, nil
}

func filterXMLXPath(content []byte, filter string) ([]byte, error) {
	doc, err := xmlquery.Parse(bytes.NewBuffer(content))
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrParsingFailed, err)
	}

	var newContent = bytes.NewBuffer(nil)

	list := xmlquery.Find(doc, filter)
	for _, node := range list {
		newContent.WriteString(node.OutputXML(true))
	}

	return newContent.Bytes(), nil
}

func filterHTMLXPath(content []byte, filter string) ([]byte, error) {
	doc, err := htmlquery.Parse(bytes.NewBuffer(content))
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrParsingFailed, err)
	}

	var newContent = bytes.NewBuffer(nil)

	list := htmlquery.Find(doc, filter)
	for _, node := range list {
		if err := html.Render(newContent, node); err != nil {
			return nil, err
		}
	}

	return newContent.Bytes(), nil
}

func filterCSS(content []byte, selector string) ([]byte, error) {
	doc, err := html.Parse(bytes.NewBuffer(content))
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrParsingFailed, err)
	}

	s, err := cascadia.Compile(selector)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrFilterInvalid, err)
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
