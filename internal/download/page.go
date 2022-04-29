package download

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func Page(uri, targetFile string) error {

	err := ensureTargetDir(targetFile)
	if err != nil {
		return fmt.Errorf("error creating target dir for file '%s': %w", targetFile, err)
	}

	resp, err := http.Get(uri)
	if err != nil {
		return err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("unexpected status code %v", resp.StatusCode)
	}
	defer func() {
		errClose := resp.Body.Close()
		if errClose != nil {
			log.Printf("error closing response body: %v", errClose)
		}
	}()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(targetFile, content, 0655)
	if err != nil {
		return fmt.Errorf("error writing page '%s' to file '%s': %w", uri, targetFile, err)
	}

	var statusFile = fmt.Sprintf("%s.response.json", targetFile)
	f, err := os.OpenFile(statusFile, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0655)
	if err != nil {
		return fmt.Errorf("error creating page '%s' status to file '%s': %w", uri, statusFile, err)
	}

	defer func() {
		errClose := f.Close()
		if errClose != nil {
			log.Printf("error closing status file '%s': %v", statusFile, errClose)
		}
	}()

	err = json.NewEncoder(f).Encode(struct {
		StatusCode       int
		Status           string
		Header           http.Header
		ContentLength    int64
		TransferEncoding []string
		DateTime         int64
	}{
		DateTime:         time.Now().UnixNano(),
		StatusCode:       resp.StatusCode,
		Status:           resp.Status,
		Header:           resp.Header,
		ContentLength:    resp.ContentLength,
		TransferEncoding: resp.TransferEncoding,
	})
	if err != nil {
		return fmt.Errorf("error writing page '%s' status to file '%s': %w", uri, statusFile, err)
	}

	return nil
}

func ensureTargetDir(file string) error {
	d := filepath.Dir(file)

	return os.MkdirAll(d, 0755)
}
