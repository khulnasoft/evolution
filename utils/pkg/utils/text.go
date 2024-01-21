package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type TextEditFunc func(string) (string, error)

func ReplaceTextTags(text, startTag, endTag, replace string) (string, error) {
	start := 0
	for {
		start = strings.Index(text[start:], startTag)
		if start < 0 {
			return text, nil
		}
		start += len(startTag)
		end := strings.Index(text[start:], endTag)
		if end < 0 {
			return "", fmt.Errorf("can't find end tag: " + endTag)
		}
		end += start
		text = text[:start] + replace + text[end:]
		start += len(replace) + len(endTag)
	}
}

func EditTextFile(path string, editors ...TextEditFunc) error {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	edited := string(bytes)
	for _, e := range editors {
		edited, err = e(edited)
		if err != nil {
			return err
		}
	}
	return ioutil.WriteFile(path, ([]byte)(edited), 0)
}

func EditCreateTextFile(path string, editors ...TextEditFunc) error {
	err := EditTextFile(path, editors...)
	if err != nil && os.IsNotExist(err) {
		var f *os.File
		f, err = os.Create(path)
		if err == nil {
			err = f.Close()
			if err == nil {
				err = EditTextFile(path, editors...)
			}
		}
	}
	return err
}
