package lib

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func LoadJsonFrom(location string) []byte {
	file, err := os.Open(filepath.Join(location))
	defer file.Close()

	if err != nil {
		log.Fatal("Error reading config file:" + err.Error())
	}

	fileString, _ := ioutil.ReadAll(file)
	return fileString
}
