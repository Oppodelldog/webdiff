package diff

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"io/ioutil"
)

func Diff(fileA, fileB string) (string, error) {
	var (
		contentA, errA = ioutil.ReadFile(fileA)
		contentB, errB = ioutil.ReadFile(fileB)
	)

	if errA != nil {
		return "", fmt.Errorf("error reading fileA '%s': %v", fileA, errA)
	}

	if errB != nil {
		return "", fmt.Errorf("error reading fileB '%s': %v", fileB, errB)
	}

	return cmp.Diff(contentA, contentB), nil
}
