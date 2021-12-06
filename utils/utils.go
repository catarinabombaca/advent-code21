package utils

import (
	"fmt"
	"io/ioutil"
)

func GetFileContent(filename string) []byte {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Errorf("error reading file: %v", err)
	}

	return b
}
