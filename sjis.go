package enc

import (
	"io/ioutil"
	"strings"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

func FromSJIS(src string) (string, error) {
	if v, err := ioutil.ReadAll(transform.NewReader(strings.NewReader(src), japanese.ShiftJIS.NewDecoder())); err != nil {
		return "", err
	} else {
		return string(v), nil
	}
}

func ToSJIS(src string) (string, error) {
	if v, err := ioutil.ReadAll(transform.NewReader(strings.NewReader(src), japanese.ShiftJIS.NewEncoder())); err != nil {
		return "", err
	} else {
		return string(v), nil
	}
}
