package enc

import (
	"io/ioutil"
	"strings"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

func FromEUCJP(src string) (string, error) {
	if v, err := ioutil.ReadAll(transform.NewReader(strings.NewReader(src), japanese.EUCJP.NewDecoder())); err != nil {
		return "", err
	} else {
		return string(v), nil
	}
}

func ToEUCJP(src string) (string, error) {
	if v, err := ioutil.ReadAll(transform.NewReader(strings.NewReader(src), japanese.EUCJP.NewEncoder())); err != nil {
		return "", err
	} else {
		return string(v), nil
	}
}
